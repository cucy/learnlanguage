
# hmac 加密

```python


import base64
import hashlib
import hmac
import json
import time
import logging
import requests

crm_server = 'https://crm.bytedance.com'
pull_clues_api = '/crm/v2/openapi/pull-clues/'


class Signature:
    """ 签名类"""

    def __init__(self, signature_header=None, timestamp_header=None, timeout=None, token_header=None, digestmod=None):
        """
        
        :param signature_header:    签名头部
        :param timestamp_header:    时间戳
        :param timeout:             超时
        :param token_header:        token
        :param digestmod:           加密算法
        """
        self.signature_header = signature_header or 'Signature'
        self.timestamp_header = timestamp_header or 'Timestamp'
        self.token_header = token_header or 'Access-Token'
        self.timeout = timeout or 300
        self.digestmod = digestmod or hashlib.sha256
        self.signature_key = bytes('456', 'utf-8')
        self.token = '123'

    def generate_signature(self, data, key=None):
        key = key if key else self.signature_key

        # signature = hmac.new(key, data, digestmod=self.digestmod).hexdigest()
        # print(type(data))
        # print(type(key))
        # key = bytes(key, 'utf-8')
        data = data.encode('utf-8')
        signature = hmac.new(key, data, digestmod=self.digestmod).hexdigest()
        print(type(signature))
        signature = base64.b64encode(bytes(signature, 'utf-8'))
        return signature


# pagepage_size, page_size10
# start_timeend_time
def crm_pull_clues(start_time, end_time, page=1, page_size=10):
    payload = {}

    payload["start_time"] = start_time
    payload["end_time"] = end_time
    payload["page"] = page
    payload["page_size"] = page_size
    timestamp = str(int(time.time()))
    sig = Signature()
    url = crm_server + pull_clues_api
    data = sig.generate_signature(
        '/crm/v2/openapi/pull-clues/?start_time=%s&end_time=%s %s' % (start_time, end_time, timestamp) )
    headers = {sig.signature_header: data, sig.timestamp_header: timestamp, sig.token_header: sig.token}
    headers.update({'Content-Type': 'application/json'})
    max_tries = 3  # 重试次数
    resp_data = None
    while max_tries > 0:
        try:
            # 打印请求信息
            t = {'payload': payload, 'headers': headers}
            print(t, '\n')
            resp = requests.get(url=url, params=payload, headers=headers)
            if resp.status_code == 200:
                res = json.loads(resp.content)  #
                if res:
                    return res
        except Exception as e:
            logging.exception("error with %s" % str(e))
            resp_data = None
        max_tries -= 1
    return resp_data


if __name__ == '__main__':

    res = crm_pull_clues("2018-10-01", "2018-10-30")
    # 打印返回值
    print('\n', '打印返回值')
    print(res)

    if res.get("status") == 0:
        data = res.get("data") or []
        if res.get["count"] > 10:
            count = res.get["count"]
            times = count / 10
            for i in range(times):  # 10
                res = crm_pull_clues("2018-07-01", "2018-07-30", 2 + i)
                data += res.get('data')
        print(len(data))
        print(data)


```
