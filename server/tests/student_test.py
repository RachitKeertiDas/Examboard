import os
import json
import openpyxl

print("Starting data extraction ...")

workbook = openpyxl.load_workbook('Students.xlsx',read_only=True)#Open Workbook

sheet = workbook.active

print("Accessed Sheet Successfully")

print(sheet['D4'].value)

row_count = sheet.max_row

students_array_dict = {}

students_array = []

current_student = sheet['B2'].value

for i in range(2,row_count+1):
	if sheet['B' + str(i)].value is None:
		continue
	else:
		print(sheet['B'+str(i)].value)#Debug
		students_dict = {}
		students_dict["Name"] = sheet['C'+str(i)].value
		students_dict["Roll No"] = sheet['B'+str(i)].value
		print(sheet['C'+str(i)].value)#debug
		courses_array = [sheet['E'+str(i)].value]
		i = i+1
		while sheet['B' + str(i)].value is None and sheet['E'+str(i)].value is not None:
			courses_array.append(sheet['E'+str(i)].value)
			i=i+1
			#print(i)
		students_dict["Courses"] = courses_array
		students_array.append(students_dict)

students_array_dict["Students"] = students_array




with open("students/students.json","w") as outfile:
	json.dump(students_array_dict,outfile,sort_keys=True,indent=4)