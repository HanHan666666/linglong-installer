#!/usr/bin/env node

const childProcess = require('child_process');
const https = require('https');
const http = require('http');
const fs = require('fs');
const path = require('path');

const CONFIG = {
  GITHUB_TOKEN: process.env.GITHUB_TOKEN || '',
  GITEE_TOKEN: process.env.GITEE_TOKEN || '',
  GITEE_USERNAME: process.env.GITEE_USERNAME || '',
  GITHUB_REPO: process.env.GITHUB_REPO || 'HanHan666666/linglong-installer',
  GITEE_REPO: process.env.GITEE_REPO || '',
  TARGET_TAG: process.env.TARGET_TAG || '',
  TRIGGER_WORKFLOW: process.env.TRIGGER_WORKFLOW !== 'false',
  WORKFLOW_FILE: process.env.WORKFLOW_FILE || 'build.yml',
  WORKFLOW_REF: process.env.WORKFLOW_REF || 'main',
  GO_VERSION: process.env.GO_VERSION || '1.24.11',
  CREATE_RELEASE: process.env.CREATE_RELEASE !== 'false',
  POLL_INTERVAL_MS: Number(process.env.POLL_INTERVAL_MS || 10000),
  WORKFLOW_TIMEOUT_MS: Number(process.env.WORKFLOW_TIMEOUT_MS || 2 * 60 * 60 * 1000),
  RELEASE_TIMEOUT_MS: Number(process.env.RELEASE_TIMEOUT_MS || 10 * 60 * 1000),
  TEMP_DIR: path.join(__dirname, '.sync-temp'),
  MAX_FILE_SIZE: 100 * 1024 * 1024,
};

function normalizeRepoPath(repo, expectedHost) {
  if (!repo) {
    return '';
  }

  const trimmed = repo.trim();
  if (!trimmed) {
    return '';
  }

  if (!trimmed.includes('://')) {
    return trimmed.replace(/^\/+|\/+$/g, '').replace(/\.git$/i, '');
  }

  const url = new URL(trimmed);
  if (expectedHost && url.hostname !== expectedHost) {
    throw new Error(`Repository URL host mismatch: expected ${expectedHost}, got ${url.hostname}`);
  }

  return url.pathname.replace(/^\/+|\/+$/g, '').replace(/\.git$/i, '');
}

CONFIG.GITHUB_REPO = normalizeRepoPath(CONFIG.GITHUB_REPO, 'github.com');
CONFIG.GITEE_REPO = normalizeRepoPath(CONFIG.GITEE_REPO, 'gitee.com');

function validateConfig() {
  const missing = [];

  if (!CONFIG.GITHUB_TOKEN) {
    missing.push('GITHUB_TOKEN');
  }
  if (!CONFIG.GITEE_TOKEN) {
    missing.push('GITEE_TOKEN');
  }
  if (!CONFIG.GITEE_REPO) {
    missing.push('GITEE_REPO');
  }
  if (CONFIG.TRIGGER_WORKFLOW && !CONFIG.TARGET_TAG) {
    missing.push('TARGET_TAG');
  }

  if (missing.length) {
    throw new Error(`Missing required configuration: ${missing.join(', ')}`);
  }
}

function request(url, options = {}) {
  return new Promise((resolve, reject) => {
    const protocol = url.startsWith('https') ? https : http;
    const req = protocol.request(url, options, (res) => {
      const chunks = [];

      if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
        resolve(request(res.headers.location, options));
        return;
      }

      res.on('data', (chunk) => chunks.push(chunk));
      res.on('end', () => {
        const body = Buffer.concat(chunks);

        if (res.statusCode >= 200 && res.statusCode < 300) {
          resolve({
            statusCode: res.statusCode,
            headers: res.headers,
            body,
          });
          return;
        }

        reject(new Error(`请求失败: ${res.statusCode} - ${body.toString('utf8')}`));
      });
    });

    req.on('error', reject);

    if (options.body) {
      req.write(options.body);
    }

    req.end();
  });
}

async function githubRequest(endpoint, options = {}) {
  const response = await request(`https://api.github.com${endpoint}`, {
    ...options,
    headers: {
      'Authorization': `Bearer ${CONFIG.GITHUB_TOKEN}`,
      'Accept': 'application/vnd.github+json',
      'X-GitHub-Api-Version': '2022-11-28',
      'User-Agent': 'linglong-installer-github2gitee',
      ...options.headers,
    },
  });

  if (!response.body.length) {
    return null;
  }

  return JSON.parse(response.body.toString('utf8'));
}

async function giteeRequest(endpoint, options = {}) {
  const separator = endpoint.includes('?') ? '&' : '?';
  const url = `https://gitee.com/api/v5${endpoint}${separator}access_token=${CONFIG.GITEE_TOKEN}`;
  const response = await request(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json;charset=UTF-8',
      ...options.headers,
    },
  });

  if (!response.body.length) {
    return null;
  }

  return JSON.parse(response.body.toString('utf8'));
}

async function giteeFormRequest(endpoint, method, payload) {
  const separator = endpoint.includes('?') ? '&' : '?';
  const formBody = new URLSearchParams();

  for (const [key, value] of Object.entries(payload)) {
    if (value === undefined || value === null) {
      continue;
    }
    formBody.append(key, String(value));
  }

  const response = await request(`https://gitee.com/api/v5${endpoint}${separator}access_token=${CONFIG.GITEE_TOKEN}`, {
    method,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
    },
    body: formBody.toString(),
  });

  if (!response.body.length) {
    return null;
  }

  return JSON.parse(response.body.toString('utf8'));
}

function ensureTempDir() {
  fs.mkdirSync(CONFIG.TEMP_DIR, { recursive: true });
}

function cleanupTempDir() {
  fs.rmSync(CONFIG.TEMP_DIR, { recursive: true, force: true });
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

function redactSecrets(text) {
  if (!text) {
    return '';
  }

  let sanitized = text;
  for (const secret of [CONFIG.GITHUB_TOKEN, CONFIG.GITEE_TOKEN]) {
    if (!secret) {
      continue;
    }
    sanitized = sanitized.split(secret).join('***');
    sanitized = sanitized.split(encodeURIComponent(secret)).join('***');
  }

  return sanitized;
}

function runCommand(command, args, options = {}) {
  const result = childProcess.spawnSync(command, args, {
    encoding: 'utf8',
    stdio: ['ignore', 'pipe', 'pipe'],
    ...options,
  });

  if (result.error) {
    throw result.error;
  }

  if (result.status !== 0) {
    const commandLine = redactSecrets(`${command} ${args.join(' ')}`.trim());
    const detail = redactSecrets((result.stderr || result.stdout || '').trim());
    throw new Error(`${commandLine} 执行失败${detail ? `: ${detail}` : ''}`);
  }

  return (result.stdout || '').trim();
}

function buildReleaseBody(release) {
  const body = (release.body || '').trim();
  if (body) {
    return body;
  }

  return [
    `Release ${release.tag_name}`,
    '',
    'No release notes were provided by GitHub.',
  ].join('\n');
}

function buildGiteeReleasePayload(release, overrides = {}) {
  return {
    tag_name: release.tag_name,
    name: release.name || release.tag_name,
    body: buildReleaseBody(release),
    prerelease: release.prerelease || false,
    // Gitee 创建 release 时会尝试同步创建 tag，所以这里的 target_commitish
    // 必须已经存在于 Gitee 仓库。主流程会先镜像 GitHub refs，再走 release 同步。
    target_commitish: release.target_commitish || 'master',
    ...overrides,
  };
}

function buildGitHubGitURL() {
  return `https://x-access-token:${encodeURIComponent(CONFIG.GITHUB_TOKEN)}@github.com/${CONFIG.GITHUB_REPO}.git`;
}

function buildGiteeGitURL(username) {
  return `https://${encodeURIComponent(username)}:${encodeURIComponent(CONFIG.GITEE_TOKEN)}@gitee.com/${CONFIG.GITEE_REPO}.git`;
}

async function getGiteeUsername() {
  if (CONFIG.GITEE_USERNAME) {
    return CONFIG.GITEE_USERNAME;
  }

  const currentUser = await giteeRequest('/user');
  if (!currentUser || !currentUser.login) {
    throw new Error('无法通过 Gitee Token 获取当前用户名，无法执行代码镜像');
  }

  CONFIG.GITEE_USERNAME = currentUser.login;
  return CONFIG.GITEE_USERNAME;
}

async function mirrorGitHubRepoToGitee() {
  console.log('🚚 开始镜像 GitHub 代码到 Gitee');

  ensureTempDir();
  const mirrorDir = path.join(CONFIG.TEMP_DIR, 'repo-mirror.git');
  fs.rmSync(mirrorDir, { recursive: true, force: true });

  const giteeUsername = await getGiteeUsername();

  console.log(`  GitHub 仓库: ${CONFIG.GITHUB_REPO}`);
  console.log(`  Gitee 仓库: ${CONFIG.GITEE_REPO}`);

  // 镜像仓库使用独立的 bare clone，避免读取本地脏工作区或误用当前 remote 配置。
  runCommand('git', ['clone', '--mirror', buildGitHubGitURL(), mirrorDir]);
  runCommand('git', ['--git-dir', mirrorDir, 'remote', 'add', 'gitee', buildGiteeGitURL(giteeUsername)]);

  // 仅同步 heads 和 tags，避免把临时 refs 或宿主环境里的额外 refs 推到 Gitee。
  runCommand('git', ['--git-dir', mirrorDir, 'push', '--prune', 'gitee', 'refs/heads/*:refs/heads/*']);
  runCommand('git', ['--git-dir', mirrorDir, 'push', '--prune', 'gitee', 'refs/tags/*:refs/tags/*']);

  console.log('  ✅ GitHub 代码已镜像到 Gitee');
}

async function dispatchGitHubWorkflow() {
  console.log(`🚀 正在触发 GitHub Actions 工作流: ${CONFIG.WORKFLOW_FILE}`);
  console.log(`  分支: ${CONFIG.WORKFLOW_REF}`);
  console.log(`  Tag: ${CONFIG.TARGET_TAG}`);

  await githubRequest(`/repos/${CONFIG.GITHUB_REPO}/actions/workflows/${encodeURIComponent(CONFIG.WORKFLOW_FILE)}/dispatches`, {
    method: 'POST',
    body: JSON.stringify({
      ref: CONFIG.WORKFLOW_REF,
      inputs: {
        go_version: CONFIG.GO_VERSION,
        release_tag: CONFIG.TARGET_TAG,
        create_release: String(CONFIG.CREATE_RELEASE),
      },
    }),
    headers: {
      'Content-Type': 'application/json;charset=UTF-8',
    },
  });
}

function findMatchingWorkflowRun(runs, startedAt) {
  const cutoff = startedAt - 60 * 1000;
  const matchedRuns = runs.filter((run) => {
    if (run.event !== 'workflow_dispatch') {
      return false;
    }
    if (run.head_branch !== CONFIG.WORKFLOW_REF) {
      return false;
    }
    if (Date.parse(run.created_at) < cutoff) {
      return false;
    }
    return true;
  });

  matchedRuns.sort((a, b) => Date.parse(b.created_at) - Date.parse(a.created_at));

  const exactTitleMatch = matchedRuns.find((run) => run.display_title === CONFIG.TARGET_TAG || run.name === CONFIG.TARGET_TAG);
  return exactTitleMatch || matchedRuns[0] || null;
}

async function waitForWorkflowRun(startedAt) {
  const deadline = Date.now() + CONFIG.WORKFLOW_TIMEOUT_MS;

  while (Date.now() < deadline) {
    const response = await githubRequest(`/repos/${CONFIG.GITHUB_REPO}/actions/workflows/${encodeURIComponent(CONFIG.WORKFLOW_FILE)}/runs?event=workflow_dispatch&branch=${encodeURIComponent(CONFIG.WORKFLOW_REF)}&per_page=20`);
    const run = findMatchingWorkflowRun(response.workflow_runs || [], startedAt);

    if (run) {
      console.log(`  ✅ 已匹配到工作流运行: #${run.run_number} (ID: ${run.id})`);
      return run;
    }

    console.log('  ⏳ 等待 GitHub Actions 创建运行记录...');
    await sleep(CONFIG.POLL_INTERVAL_MS);
  }

  throw new Error(`等待工作流运行记录超时: ${CONFIG.WORKFLOW_FILE}`);
}

async function waitForWorkflowCompletion(runId) {
  const deadline = Date.now() + CONFIG.WORKFLOW_TIMEOUT_MS;

  while (Date.now() < deadline) {
    const run = await githubRequest(`/repos/${CONFIG.GITHUB_REPO}/actions/runs/${runId}`);
    console.log(`  ⏳ 工作流状态: ${run.status}${run.conclusion ? ` / ${run.conclusion}` : ''}`);

    if (run.status === 'completed') {
      if (run.conclusion !== 'success') {
        throw new Error(`GitHub Actions 构建失败: ${run.conclusion} (${run.html_url})`);
      }

      console.log(`  ✅ GitHub Actions 构建完成: ${run.html_url}`);
      return run;
    }

    await sleep(CONFIG.POLL_INTERVAL_MS);
  }

  throw new Error(`等待工作流执行完成超时: run_id=${runId}`);
}

async function waitForGitHubReleaseVisibility(tagName) {
  const deadline = Date.now() + CONFIG.RELEASE_TIMEOUT_MS;

  while (Date.now() < deadline) {
    try {
      const release = await githubRequest(`/repos/${CONFIG.GITHUB_REPO}/releases/tags/${encodeURIComponent(tagName)}`);
      console.log(`  ✅ GitHub Release 已可见: ${release.html_url}`);
      return release;
    } catch (error) {
      if (!error.message.includes('404')) {
        throw error;
      }
    }

    console.log(`  ⏳ 等待 GitHub Release ${tagName} 可见...`);
    await sleep(CONFIG.POLL_INTERVAL_MS);
  }

  throw new Error(`等待 GitHub Release 可见超时: ${tagName}`);
}

async function runWorkflowAndWaitForRelease() {
  try {
    const existingRelease = await githubRequest(`/repos/${CONFIG.GITHUB_REPO}/releases/tags/${encodeURIComponent(CONFIG.TARGET_TAG)}`);
    console.log(`📦 GitHub Release 已存在，跳过构建触发: ${existingRelease.html_url}`);
    return existingRelease;
  } catch (error) {
    if (!error.message.includes('404')) {
      throw error;
    }
  }

  const startedAt = Date.now();
  await dispatchGitHubWorkflow();
  const run = await waitForWorkflowRun(startedAt);
  await waitForWorkflowCompletion(run.id);
  return waitForGitHubReleaseVisibility(CONFIG.TARGET_TAG);
}

async function getGitHubRelease() {
  if (CONFIG.TARGET_TAG) {
    console.log(`📥 正在获取 GitHub Release: ${CONFIG.TARGET_TAG}`);
    return githubRequest(`/repos/${CONFIG.GITHUB_REPO}/releases/tags/${CONFIG.TARGET_TAG}`);
  }

  console.log('📥 正在获取 GitHub 最新 Release');
  const releases = await githubRequest(`/repos/${CONFIG.GITHUB_REPO}/releases?per_page=1`);
  if (!releases.length) {
    throw new Error('GitHub 上没有可同步的 Release');
  }
  return releases[0];
}

async function getGiteeRelease(tagName) {
  console.log(`📥 正在检查 Gitee Release: ${tagName}`);

  try {
    const releases = await giteeRequest(`/repos/${CONFIG.GITEE_REPO}/releases?page=1&per_page=100`);
    return releases.find((release) => release.tag_name === tagName) || null;
  } catch (error) {
    console.log('⚠️  获取 Gitee Release 列表失败，按不存在处理');
    return null;
  }
}

async function downloadFile(url, filepath) {
  return new Promise((resolve, reject) => {
    const protocol = url.startsWith('https') ? https : http;
    const file = fs.createWriteStream(filepath);
    const req = protocol.get(url, {
      headers: {
        'Authorization': `token ${CONFIG.GITHUB_TOKEN}`,
        'Accept': 'application/octet-stream',
        'User-Agent': 'linglong-installer-github2gitee',
      },
    }, (res) => {
      if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
        file.close();
        fs.rmSync(filepath, { force: true });
        resolve(downloadFile(res.headers.location, filepath));
        return;
      }

      if (res.statusCode !== 200) {
        file.close();
        fs.rmSync(filepath, { force: true });
        reject(new Error(`下载失败: ${res.statusCode}`));
        return;
      }

      res.pipe(file);
      file.on('finish', () => file.close(() => resolve(filepath)));
    });

    req.on('error', (error) => {
      file.close();
      fs.rmSync(filepath, { force: true });
      reject(error);
    });

    file.on('error', (error) => {
      file.close();
      fs.rmSync(filepath, { force: true });
      reject(error);
    });
  });
}

async function uploadAssetToGiteeRelease(releaseId, filepath) {
  const filename = path.basename(filepath);
  const stats = fs.statSync(filepath);
  const boundary = `----codex${Date.now()}${Math.random().toString(16).slice(2)}`;
  const header = Buffer.from([
    `--${boundary}`,
    `Content-Disposition: form-data; name="file"; filename="${filename}"`,
    'Content-Type: application/octet-stream',
    '',
    '',
  ].join('\r\n'));
  const footer = Buffer.from(`\r\n--${boundary}--\r\n`);
  const url = `https://gitee.com/api/v5/repos/${CONFIG.GITEE_REPO}/releases/${releaseId}/attach_files?access_token=${CONFIG.GITEE_TOKEN}`;
  const urlObj = new URL(url);

  return new Promise((resolve, reject) => {
    const req = https.request({
      hostname: urlObj.hostname,
      port: urlObj.port || 443,
      path: urlObj.pathname + urlObj.search,
      method: 'POST',
      headers: {
        'Content-Type': `multipart/form-data; boundary=${boundary}`,
        'Content-Length': header.length + stats.size + footer.length,
      },
    }, (res) => {
      const chunks = [];
      res.on('data', (chunk) => chunks.push(chunk));
      res.on('end', () => {
        const body = Buffer.concat(chunks).toString('utf8');
        if (res.statusCode >= 200 && res.statusCode < 300) {
          resolve(body ? JSON.parse(body) : { name: filename, size: stats.size });
          return;
        }
        reject(new Error(`上传失败 (${res.statusCode}): ${body}`));
      });
    });

    req.on('error', reject);
    req.write(header);

    const fileStream = fs.createReadStream(filepath);
    fileStream.on('error', (error) => {
      req.destroy(error);
    });
    fileStream.on('data', (chunk) => req.write(chunk));
    fileStream.on('end', () => {
      req.write(footer);
      req.end();
    });
  });
}

async function createGiteeRelease(release) {
  return giteeFormRequest(`/repos/${CONFIG.GITEE_REPO}/releases`, 'POST', buildGiteeReleasePayload(release));
}

async function updateGiteeRelease(releaseId, release, overrides = {}) {
  const payload = buildGiteeReleasePayload(release, overrides);
  delete payload.target_commitish;
  return giteeFormRequest(`/repos/${CONFIG.GITEE_REPO}/releases/${releaseId}`, 'PATCH', payload);
}

function getMissingAssets(githubRelease, giteeRelease) {
  const githubAssets = githubRelease.assets || [];
  const giteeAssets = new Map(((giteeRelease && giteeRelease.assets) || []).map((asset) => [asset.name, asset.size]));

  return githubAssets.filter((asset) => {
    const existingSize = giteeAssets.get(asset.name);
    if (existingSize === undefined) {
      return true;
    }
    if (existingSize === asset.size) {
      console.log(`  ✅ 附件 "${asset.name}" 已存在，跳过`);
      return false;
    }

    console.warn(`  ⚠️  附件 "${asset.name}" 已存在但大小不同，保留 Gitee 现有文件`);
    return false;
  });
}

async function syncRelease() {
  console.log('🚀 开始同步 GitHub 最新 Release 到 Gitee\n');
  console.log(`GitHub 仓库: ${CONFIG.GITHUB_REPO}`);
  console.log(`Gitee 仓库: ${CONFIG.GITEE_REPO}\n`);

  const githubRelease = await getGitHubRelease();
  const tagName = githubRelease.tag_name;
  let giteeRelease = await getGiteeRelease(tagName);
  const finalPrerelease = githubRelease.prerelease || false;
  const createdRelease = !giteeRelease;

  console.log(`\n📦 处理 Release: ${tagName}`);

  if (createdRelease) {
    console.log('  🚀 Gitee 不存在该 Release，准备先创建为预发布');
    giteeRelease = await createGiteeRelease({
      ...githubRelease,
      prerelease: true,
    });
    console.log(`  ✅ Release 创建成功 (ID: ${giteeRelease.id})`);
  } else {
    console.log('  ℹ️  Gitee 已存在该 Release，仅检查缺失附件');
  }

  const assets = getMissingAssets(githubRelease, giteeRelease);

  if (!assets.length) {
    console.log('  ✅ 无需同步附件');
    if (createdRelease && giteeRelease.prerelease !== finalPrerelease) {
      console.log(`  🚀 将 Release 状态恢复为${finalPrerelease ? '预发布' : '正式发布(latest)'}`);
      await updateGiteeRelease(giteeRelease.id, githubRelease, { prerelease: finalPrerelease });
      console.log('  ✅ Release 状态已更新');
    }
    return;
  }

  ensureTempDir();
  let uploadedCount = 0;
  let skippedCount = 0;

  for (const asset of assets) {
    const tempFile = path.join(CONFIG.TEMP_DIR, asset.name);

    try {
      console.log(`\n    📄 处理: ${asset.name} (${(asset.size / 1024 / 1024).toFixed(2)} MB)`);
      if (asset.size > CONFIG.MAX_FILE_SIZE) {
        console.warn(`      ⚠️  文件超过 ${CONFIG.MAX_FILE_SIZE / 1024 / 1024}MB，跳过`);
        skippedCount += 1;
        continue;
      }

      await downloadFile(asset.browser_download_url, tempFile);
      console.log('      ✓ 下载完成');
      await uploadAssetToGiteeRelease(giteeRelease.id, tempFile);
      console.log('      ✓ 上传成功');
      uploadedCount += 1;
    } finally {
      fs.rmSync(tempFile, { force: true });
    }
  }

  console.log(`\n  ✅ 附件处理完成: ${uploadedCount}/${assets.length} 个成功上传`);

  if (skippedCount > 0) {
    throw new Error(`有 ${skippedCount} 个附件因大小限制未上传`);
  }

  if (createdRelease && giteeRelease.prerelease !== finalPrerelease) {
    console.log(`  🚀 所有附件已上传，正在切换为${finalPrerelease ? '预发布' : '正式发布(latest)'}`);
    await updateGiteeRelease(giteeRelease.id, githubRelease, { prerelease: finalPrerelease });
    console.log('  ✅ Release 已切换到最终状态');
  }
}

async function main() {
  try {
    validateConfig();
    if (CONFIG.TRIGGER_WORKFLOW) {
      await runWorkflowAndWaitForRelease();
      console.log('');
    }
    await mirrorGitHubRepoToGitee();
    console.log('');
    await syncRelease();
    cleanupTempDir();
    console.log('\n✨ 同步完成！');
  } catch (error) {
    cleanupTempDir();
    console.error(`\n❌ 同步失败: ${error.message}`);
    process.exit(1);
  }
}

main();
