import json
import os
import pandas as pd

"""
all_courses = os.listdir("courses_dummy")

workbook = openpyxl.load_workbook('Students.xlsx',read_only=True)#Open Workbook

sheet = workbook.active


row_count = sheet.max_row

print(all_courses)

for i in range(0,row_count+1):
	print(sheet['E'+ str(i)].value)
	print(i)

"""
"""
for each_course in all_courses:
	print(each_course)
	course_name = each_course[:6]
	print(course_name)
	for i in range(0,row_count+1):
		print(i)
		if sheet['E'+ str(i)].value == course_name:
			print("FOUND"+str(i))
			course_file = open("courses_dummy/"+ each_course,"r")
			json_data = json.load(course_file)
			course_file.close()
			json_data["Instructor"] = sheet['H'+ str(i)].value
			course_file = open("courses_dummy/"+ each_course,"w+")
			json.dump(json_data,course_file,sort_keys=True,indent=2)
			course_file.close()
			break

already_done_courses = []

"""

df = pd.ExcelFile('Students.xlsx').parse(0) #open Excel Sheet in Pandas

course_codes =[]
instructors = []
course_names = []

course_codes.append(df['Course Code'])# Pandas df of Column of course_codes
instructors.append(df['Coordinator/Instructor Name'])#Pandas df of column of Instructors
course_names.append(df['Course Name'])# Pandas df of Column of course_instructors
print(len(course_codes))

total_rows = len(df.index)

courses_list =[]
instructors_list = []
course_name_list = []

#Convert Pandas Df into a regular python list
for i in range(0,total_rows):
	for each in course_codes:
		courses_list.append(each[i])

#Convert Pandas df into a regular python list
for i in range(0,total_rows):
	for each in instructors:
		instructors_list.append(each[i])

#Convert Pandas df into a regular python list
for i in range(0,total_rows):
	for each in course_names:
		course_name_list.append(each[i])

print(course_name_list)


#
#zipped_list  = list(zip(courses_list,instructors_list))

courses_dict = {}

course_name_dict = {}

for i in range(0,len(instructors_list)):#dictionary of courses
	courses_dict[str(courses_list[i])] = str(instructors_list[i])
	course_name_dict[str(courses_list[i])] = str(course_name_list[i])
#print(courses_dict)


all_courses = os.listdir("../courses")

for each_course in all_courses: #each_course is a file name
	print(each_course)
	course_code = each_course[:6]
	course_instructor = courses_dict[str(course_code)]
	course_name = course_name_dict[str(course_code)]
	course_file = open("../courses/"+ each_course,"r")
	json_data = json.load(course_file)
	course_file.close()
	json_data["Instructor"] = course_instructor
	json_data["Name"] = course_name
	course_file = open("../courses/"+ each_course,"w+")
	json.dump(json_data,course_file,sort_keys=True,indent=2)
	course_file.close()
