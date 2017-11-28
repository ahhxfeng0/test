#！/usr/bin/env/python
import os
import io
import sqlite3
import itertools
from flask import Flask
from flask import render_template
app = Flask(__name__)

DBNAME = 'init.sql'
DBUSER = 'root'
DB_EXC = None

def db_connect(db):
	global DB_EXC
	dbdir = '%s_%s' % (db, DBNAME)
	DB_EXC = sqlite3
	if not os.path.isdir(dbdir):
		os.mkdir(dbdir)
	cxn = sqlite3.connect(os.path.join(dbdir, DBNAME))
def get_conn(path):
	# self.path = path
	conn = sqlite3.connect(path, check_same_thread = False)
	return conn
# def get_cursor(conn):
# 	if conn is not None:
# 		return conn.cursor()
# 	else:
# 		return get_conn('').cursor()
def creat(conn):
	cur = conn.cursor()
	try:
		cur.execute('''
			CREATE TABLE data(
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
		''')
	finally:
		conn.close()
# except DB_EXC.OperationalError as e:

def insert_info(conn):
	cur = conn.cursor()
	for dict_mock in datas1:
		print('插入一条数据')
		# for key in keys:
		# 	print(dict_mock[key])
		cur.execute('INSERT INTO data VALUES(\
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)', [(dict_mock[key]) for key in keys])
	conn.commit()

# def select(conn):

def select_info(conn):
    entries = []
    cur = conn.cursor()
    try:
        cur.execute('SELECT * FROM data')
        while True:
            entry = cur.fetchone()
            print(entry)
            if entry is None:
                break
            entries.append(dict(zip(keys, entry)))
    finally:
        conn.close()
    return entries

	
		

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
	conn = get_conn(DBNAME)
	# creat(conn)
	open_data()
	insert_info(conn)
	datas = select_info(conn)
	return render_template('show.html', list = datas, list1 = keys)
@app.route('/configuration')
# def configuation(list):
	# datas = select_info(conn)	
# 	return render_template('configuration.html', list = datas)
@app.route('/hello')
def hello(name = None):
	return render_template('hello.html', name = name)

if __name__ == '__main__':
	app.run(debug = True)
