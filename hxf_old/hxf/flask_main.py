import os
import io
import itertools
from flask import Flask
from flask import render_template

keys = ['id', 'uuid', 'name', 'configuration', 'model_path', 'rank', 'trueskill_args', 'status', 'committer', 'commit_time', 'last_processed_time']
datas = []
def open_data():
	fp = open('mock.data', 'r+')
	for eachLine in fp:
		dict_mock = dict()
		for key, value in itertools.zip_longest(keys, eachLine.split('\t')):
			dict_mock[key] = value
		datas.append(dict_mock)
	fp.close()

@app.route('/')
def show(list = datas, list1 = keys):
	return render_template('show.html', list = list, list1 = list1)
@app.route('/configuration')
def configuation(list = datas):
	return render_template('configuration.html', list = list)
@app.route('/hello')
def hello(name = None):
	return render_template('hello.html', name = name)

if __name__ == '__main__':
	app = Flask(__name__)
	open_data()
	app.run(debug = True)
