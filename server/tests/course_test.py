import json
import os

json_data ={}

with open("students/students.json") as inp:
	json_data = json.load(inp)

count=0

for each_student in json_data['Students']:
	#print(each_student['Courses'])
	count = count+1
	for each_course in each_student['Courses']:
		course_file = open('courses/' + each_course +'.json','a')
		course_file.close()
		course_file = open('courses/' + each_course +'.json','r')		
		course_file = open('courses/' + each_course +'.json','r')
		course_data={}
		try:
			course_data = json.load(course_file)
			print("Never came here")
		except:
			print("Here")
			course_data = {}
			course_data['Students'] =[]
			course_data['Conflicts']=[]
		
		course_file.close()

		#print(course_data)
		course_data['Code'] = each_course
		course_data['Students'].append(each_student['RollNo'])
		#print(course_data['Students'])
		for each_item in each_student['Courses']:
			if each_item != each_course :
				if each_item in course_data['Conflicts']:
					continue
				course_data['Conflicts'].append(each_item)
			else:
				continue
		#print(course_data)
		
		course_file = open('courses/' + each_course +'.json','w')
		json.dump(course_data,course_file,sort_keys=True,indent=2)
		course_file.close()
		#for each_course in each_student[Courses]:
		#	print(each_student)
print(count)

#print(json_data)