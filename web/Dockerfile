FROM nginx:1.27-alpine
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV TZ=Asia/Shanghai
COPY nginx.conf /etc/nginx/nginx.conf
COPY dist /usr/share/nginx/html
