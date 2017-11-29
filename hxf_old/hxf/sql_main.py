import sqlite3
from flask import g
import os
import io
import itertools
from flask import Flask
from flask import render_template

DATABASE = 'init.db'

def connect_db():
	return sqlite3.connect(DATABASE)
@app.before_request
def before_request():
	g.db = connect_db()
@app.teardown_request
def teardown_request(Exception):
	if hasattr(g, 'db'):
		g.db.close()
