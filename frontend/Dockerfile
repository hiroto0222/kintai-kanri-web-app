FROM node:18-alpine3.17

WORKDIR /app
COPY ./package.json /app
COPY ./yarn.lock /app

# yarn install だと依存パッケージのバージョンが固定されない
# --frozen-lockfile -> yarn.lock ファイルを生成せず、更新が必要な場合は失敗する
RUN yarn install --frozen-lockfile

COPY . /app

EXPOSE 5173

CMD ["yarn", "dev"]
