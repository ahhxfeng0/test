#！/usr/bin/env/python
import os
import io
import sqlite3
import itertools
import db
from flask import Flask
from flask import render_template
app = Flask(__name__)

keys = ['id', 'uuid', 'name', 'configuration', 'model_path', 'rank', 'trueskill_args', 'status', 'committer', 'commit_time', 'last_processed_time']
datas1 = []
def open_data():
	fp = open('mock.data', 'r+')
	for eachLine in fp:
		dict_mock = dict()
		for key, value in itertools.zip_longest(keys, eachLine.split('\t')):
			dict_mock[key] = value
		datas1.append(dict_mock)
	fp.close()



@app.route('/')
def show():
	conn = db.get_conn(db.DBNAME)
	db.creat(conn, db.table_name)
	open_data()
	db.insert_info(conn)
	datas = db.select_info(conn, db.table_name)
	return render_template('show.html', list = datas, list1 = keys)
@app.route('/configuration')
def configuation():
	conn = db.get_conn(db.DBNAME)
	datas = db.select_info(conn, db.table_name)
	return render_template('configuration.html', list = datas)
@app.route('/hello')
def hello(name = None):
	return render_template('hello.html', name = name)

if __name__ == '__main__':
	app.run(debug = True)
