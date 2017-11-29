import os
import io
import sqlite3
import flask_main
import itertools
from flask import Flask
from flask import render_template

DBNAME = 'init.sql'
table_name = "data"
keys = flask_main.keys
datas1 = flask_main.datas1

def get_conn(path):
	# self.path = path
	conn = sqlite3.connect(path, check_same_thread = False)
	return conn

def creat(conn, table_name):
	cur = conn.cursor()
	try:
		print('创建数据库')
		cur.execute('''
			CREATE TABLE IF NOT EXISTS %s(
				'id' INTEGER,
				'uuid' TEXT,
				'name' TEXT,
				'configuration' TEXT,
				'model_path' TEXT,
				'rank' REAL,
				'trueskill_args' TEXT,
				'status' TEXT,
				'committer' TEXT,
				'commit_time' TEXT,
				'last_processed_time' TEXT)
		''' % table_name)
	finally:
		conn.commit()

def insert_info(conn):
	flask_main.open_data()
	cur = conn.cursor()
	print('插入数据')
	for dict_mock in datas1:
		print('插入一条数据')
		for key in keys:
			print(dict_mock[key])
		cur.execute('INSERT INTO data VALUES(\
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)', [(dict_mock[key]) for key in keys])
	conn.commit()


def select_info(conn, table_name):
    datas = []
    cur = conn.cursor()
    result = cur.execute('SELECT * FROM %s' % table_name)
    for eachLine in result:
        data = dict()
        for key, value in zip(keys, eachLine):
            data[key] = value
        datas.append(data)
    return datas

