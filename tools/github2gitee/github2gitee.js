#!/usr/bin/env node

const https = require('https');
const http = require('http');
const fs = require('fs');
const path = require('path');

const CONFIG = {
  GITHUB_TOKEN: process.env.GITHUB_TOKEN || '',
  GITEE_TOKEN: process.env.GITEE_TOKEN || '',
  GITHUB_REPO: process.env.GITHUB_REPO || 'HanHan666666/linglong-installer',
  GITEE_REPO: process.env.GITEE_REPO || '',
  TARGET_TAG: process.env.TARGET_TAG || '',
  TEMP_DIR: path.join(__dirname, '.sync-temp'),
  MAX_FILE_SIZE: 100 * 1024 * 1024,
};

if (!CONFIG.GITHUB_TOKEN || !CONFIG.GITEE_TOKEN || !CONFIG.GITEE_REPO) {
  console.error('Missing required configuration: GITHUB_TOKEN, GITEE_TOKEN, GITEE_REPO');
  process.exit(1);
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
      'Authorization': `token ${CONFIG.GITHUB_TOKEN}`,
      'Accept': 'application/vnd.github.v3+json',
      'User-Agent': 'linglong-installer-github2gitee',
      ...options.headers,
    },
  });

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

function ensureTempDir() {
  fs.mkdirSync(CONFIG.TEMP_DIR, { recursive: true });
}

function cleanupTempDir() {
  fs.rmSync(CONFIG.TEMP_DIR, { recursive: true, force: true });
}

function formatBytes(bytes) {
  if (!Number.isFinite(bytes) || bytes <= 0) {
    return '0 B';
  }

  const units = ['B', 'KB', 'MB', 'GB', 'TB'];
  const exponent = Math.min(Math.floor(Math.log(bytes) / Math.log(1024)), units.length - 1);
  const value = bytes / (1024 ** exponent);
  const digits = value >= 100 || exponent === 0 ? 0 : value >= 10 ? 1 : 2;

  return `${value.toFixed(digits)} ${units[exponent]}`;
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
  const totalBytes = stats.size;

  return new Promise((resolve, reject) => {
    const startedAt = Date.now();
    let uploadedBytes = 0;
    let settled = false;
    let progressTimer = null;

    const printProgress = (force = false) => {
      const elapsedSeconds = Math.max((Date.now() - startedAt) / 1000, 0.001);
      const percentage = totalBytes === 0 ? 100 : (uploadedBytes / totalBytes) * 100;
      const speed = uploadedBytes / elapsedSeconds;
      const line = `      ↑ 上传进度: ${percentage.toFixed(1)}% (${formatBytes(uploadedBytes)}/${formatBytes(totalBytes)}) ${formatBytes(speed)}/s`;

      if (force) {
        process.stdout.write(`${line}\n`);
        return;
      }

      process.stdout.write(`\r${line}`);
    };

    const finish = (handler, value) => {
      if (settled) {
        return;
      }

      settled = true;
      if (progressTimer) {
        clearInterval(progressTimer);
      }

      if (uploadedBytes > 0 || totalBytes === 0) {
        printProgress(true);
      }

      handler(value);
    };

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
          finish(resolve, body ? JSON.parse(body) : { name: filename, size: stats.size });
          return;
        }
        finish(reject, new Error(`上传失败 (${res.statusCode}): ${body}`));
      });
    });

    req.on('error', (error) => finish(reject, error));
    req.write(header);
    progressTimer = setInterval(() => printProgress(false), 1000);

    const fileStream = fs.createReadStream(filepath);
    fileStream.on('error', (error) => {
      req.destroy(error);
    });
    fileStream.on('data', (chunk) => {
      uploadedBytes += chunk.length;

      if (!req.write(chunk)) {
        fileStream.pause();
      }
    });
    req.on('drain', () => {
      fileStream.resume();
    });
    fileStream.on('end', () => {
      req.write(footer);
      req.end();
    });
  });
}

async function createGiteeRelease(release) {
  const payload = {
    tag_name: release.tag_name,
    name: release.name || release.tag_name,
    body: buildReleaseBody(release),
    prerelease: release.prerelease || false,
    target_commitish: release.target_commitish || 'master',
  };

  return giteeRequest(`/repos/${CONFIG.GITEE_REPO}/releases`, {
    method: 'POST',
    body: JSON.stringify(payload),
  });
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

  console.log(`\n📦 处理 Release: ${tagName}`);

  if (!giteeRelease) {
    console.log('  🚀 Gitee 不存在该 Release，准备创建');
    giteeRelease = await createGiteeRelease(githubRelease);
    console.log(`  ✅ Release 创建成功 (ID: ${giteeRelease.id})`);
  } else {
    console.log('  ℹ️  Gitee 已存在该 Release，仅检查缺失附件');
  }

  const assets = getMissingAssets(githubRelease, giteeRelease);
  if (!assets.length) {
    console.log('  ✅ 无需同步附件');
    return;
  }

  ensureTempDir();
  let uploadedCount = 0;

  for (const asset of assets) {
    const tempFile = path.join(CONFIG.TEMP_DIR, asset.name);

    try {
      console.log(`\n    📄 处理: ${asset.name} (${(asset.size / 1024 / 1024).toFixed(2)} MB)`);
      if (asset.size > CONFIG.MAX_FILE_SIZE) {
        console.warn(`      ⚠️  文件超过 ${CONFIG.MAX_FILE_SIZE / 1024 / 1024}MB，跳过`);
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
}

async function main() {
  try {
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
