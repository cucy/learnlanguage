import json
from protos import routeguide_pb2

"""
此文件类似数据库功能:
做加载数据用,无其他特别一样
"""


def read_route_guide_database():
    """
    Reads the route guide database. 从存储中获取数据

    Returns:
        The full contents of the route guide database as a sequence of  route_guide_pb2.Features.
    """

    feature_list = []
    with open("route_guide_db.json", encoding='utf-8') as route_guide_db_file:
        for item in json.load(route_guide_db_file, ):
            feature = routeguide_pb2.Feature(
                name=item["name"],
                location=routeguide_pb2.Point(
                    latitude=item["location"]["latitude"],
                    longitude=item["location"]["longitude"]))
            feature_list.append(feature)
    return feature_list


if __name__ == '__main__':
    print(read_route_guide_database()[0])
