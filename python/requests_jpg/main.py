import requests

from util import log


def download_jpg(url):
    # 发送HTTP GET请求
    response = requests.get(url)

    # 检查响应状态码
    if response.status_code == 200:
        # 获取文件名
        file_name = url.split("/")[-1]

        # 保存图片
        with open("assets/" + file_name, "wb") as file:
            file.write(response.content)
            log.info(f"{file_name} 图片已保存")
    else:
        log.info(f"下载图片失败:{url}")


def url_list():
    with open('list.txt', 'r', newline='\n') as f:
        urls = f.readlines()
        urls = [line.rstrip('\n') for line in urls]

        urls=set(urls)
        log.info(f'urls总数: {len(urls)}')
    # print(urls)

    for url in urls:
        log.info(f"下载url:{url}")
        download_jpg(url)



if __name__ == '__main__':
    url_list()
