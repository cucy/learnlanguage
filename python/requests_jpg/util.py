# coding=utf-8
# 把内容写入csv文件的方法
import csv
import datetime
import logging
import os
import time
from configparser import ConfigParser
from pathlib import Path



def save_search_for_product_html(content):
    save_file(content, cfg('save_search_for_product_html_abs_path'), 'w')


def save_search_lists_of_companies_to_html(content):
    file_path = cfg('search_lists_of_companies_html_abs_path')
    save_file(content, file_path, 'w')


def save_company_profile_to_html(content):
    file_path = cfg('company_profile_url_abs_path')
    save_file(content, file_path, 'w')


def read_company_profile():
    file_path = cfg('company_profile_url_abs_path')
    return read_file(file_path)


def save_file(content, file_path, mode):
    current_path = Path().resolve()
    file_abs_path = os.path.join(current_path, file_path)
    log.info(f"保存文件:{file_abs_path}")
    with open(file_abs_path, mode, encoding='utf-8') as f:
        f.write(content)


def save_file_lines(content, file_path, mode):
    current_path = Path().resolve()
    file_abs_path = os.path.join(current_path, file_path)
    log.info(f"保存文件:{file_abs_path}")
    with open(file_abs_path, mode, encoding='utf-8') as f:
        f.writelines(content)


def read_file_to_list(file_path):
    with open(file_path, 'r', encoding='utf-8') as f:
        return [i.strip() for i in f.readlines()]


def read_file(file_path):
    current_path = Path().resolve()
    file_abs_path = os.path.join(current_path, file_path)
    log.info(f"读取文件:{file_abs_path}")

    with open(file_path, 'r', encoding='utf-8') as f:
        return f.read()


def save_search_for_product_urls(site_url: list, ):
    file_abs_path = os.path.join(Path().resolve(), "search_for_product_list.csv")
    csvf = open(file_abs_path, 'a', newline='', encoding='utf-8')  # 覆盖写 ,# newline=''去除空格
    writer = csv.writer(csvf)
    size = os.path.getsize(file_abs_path)
    if size == 0:
        # headfile = ['url']
        # writer.writerow(headfile)
        writer.writerow(site_url)
        csvf.close()
    else:
        writer.writerow(site_url)
        csvf.close()


def get_company_urls_from_file() -> list[str]:
    """
    从文件读取URL列表
    :return:
    """
    return read_file_to_list(cfg('company_url_list_abs_path'))




def app_log():
    logger = logging.getLogger(__name__)
    logger.setLevel(level=logging.DEBUG)

    logging.basicConfig(level=logging.INFO,
                        format='%(asctime)s -%(filename)s- %(lineno)d - %(levelname)s - %(message)s')

    # StreamHandler
    # stream_handler = logging.StreamHandler(sys.stdout)
    # stream_handler.setLevel(level=logging.DEBUG)
    # logger.addHandler(stream_handler)
    return logger



def data_save_format(keyword: str, save_type):
    """

    :param keyword:
    :param save_type:   products | company
    :return:
    """
    now = datetime.datetime.now()
    formatted_datetime = now.strftime("%Y%m%d")
    keyword = str(keyword)
    keyword = keyword.replace(" ", "_")
    keyword_strip = keyword
    keyword_strip = keyword_strip.replace(" ", "_")

    # 最终文件保存路径
    s_path = f"{formatted_datetime}_{keyword_strip}_{save_type}_list.json"
    return s_path


def setup():
    """
    添加数据目录
    :return:
    """
    current_path = Path().resolve()

    try:
        dirs = [cfg("save_base_dir"), cfg("data_dir"), cfg("fix_dir")]
        for clear_dir in dirs:
            if os.path.exists(clear_dir):
                log.info(f"已经存在目录:{clear_dir}")
                continue
            else:
                create_dir = os.path.join(current_path, clear_dir)
                log.info(f"创建目录:{create_dir}")
                os.mkdir(create_dir)
    except Exception as e:
        pass


def cfg(key):
    """
    读取配置文件
    :return:
    """
    current_path = Path().resolve()
    cfg_abs_path = os.path.join(current_path, 'config.ini')
    cf = ConfigParser()
    cf.read(cfg_abs_path, encoding='utf-8')
    return cf.get('common', key)


log = app_log()
if __name__ == '__main__':
    a = cfg('product_count')
    print(a)
