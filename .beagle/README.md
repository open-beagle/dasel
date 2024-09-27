# dasel

<https://github.com/TomWright/dasel>

```bash
git remote add upstream git@github.com:TomWright/dasel.git

git fetch upstream

git merge v2.8.1
```

## build

```bash
# cross
docker run --rm -it \
--rm \
-v $PWD/:/go/src/github.com/TomWright/dasel \
-w /go/src/github.com/TomWright/dasel \
-e BUILD_VERSION=v2.8.1 \
registry.cn-qingdao.aliyuncs.com/wod/golang:1.22 \
bash .beagle/build.sh

# loong64
docker run --rm -it \
--rm \
-v $PWD/:/go/src/github.com/TomWright/dasel \
-w /go/src/github.com/TomWright/dasel \
-e BUILD_VERSION=v2.8.1 \
registry.cn-qingdao.aliyuncs.com/wod/golang:1.22-loongnix \
bash .beagle/build.sh
```

## cache

```bash
# 构建缓存-->推送缓存至服务器
docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="dasel" \
  -e PLUGIN_MOUNT="./.git" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0

# 读取缓存-->将缓存从服务器拉取到本地
docker run --rm \
  -e PLUGIN_RESTORE=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="dasel" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0
```
