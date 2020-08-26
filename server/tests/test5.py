import os
import time
import json

course_file_list = os.listdir("courses/")

for each_course in course_file_list:
	course_file = open(each_course,"r+")
	course_data = json.load(course_file)
"""

Expected JSON Strcuture:

Name : "Abhinav Kumar"
Courses :[{Code:"AI1102","Probability"},{Code:"EE1370","Name:Data Analytics"}]

"""	
