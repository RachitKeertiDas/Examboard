import React, { useState, useEffect } from "react";
import { Heading, Flex, Box, Divider, Text } from "@chakra-ui/core";
import "./styles.css";

function StudentView() {
  //example as seen from react website
  const [Error, SetError] = useState(null);
  const [StudentName, SetStudentName] = useState(null);
  const [Courses, SetCourses] = useState([]);
  const [RollNo, SetRollNo] = useState();
  const [isLoaded, SetisLoaded] = useState(false);

  useEffect(() => {
    if (!Courses.length) {
      fetch("http://localhost:8080/api/student/getCourses/CS19BTECH11034")
        .then((res) => res.json())
        .then(
          (result) => {
            SetisLoaded(true);
            SetStudentName(result.Name);
            SetCourses(result.Courses);
            SetRollNo(result.RollNo);
            console.log(StudentName);
            console.log(RollNo);
            console.log(Courses.length);
          },
          (error) => {
            SetisLoaded(true);
            SetError(error);
          }
        );
    }
  }, [StudentName, RollNo, Courses]);
  /*
  let StudentJSON = fetch("http://localhost:8080/api/student/getCourses/");
  let StudentJSONData;
  StudentJSON.then((response) => response.json())
    .then((StudentJSONBody) => {
      console.log("Body Recieved");
      //console.log(StudentJSONBody.Name);
      StudentJSONData = StudentJSONBody;
      console.log(StudentJSONData.Name);
    })
    .catch((error) => {
      console.log(error);
    });
  */
  return (
    <Flex
      flexDirection="column"
      justifyContent="space-between"
      alignItems="center"
    >
      <Box className="studentview-heading">
        <Heading>Hello, {StudentName}</Heading>
        <Divider />
      </Box>
      <Box>
        <Heading size="md">
          Your Courses for this semester:
          <Divider />
        </Heading>
        <table class="courses-table">
          <thead>
            <tr>
              <th>Course Code</th>
              <th>Instructor</th>
              <th>Exam Dates</th>
            </tr>
          </thead>
          <tbody>
            {Courses.map((item) => {
              return (
                <tr>
                  <td>
                    <Text>{item}</Text>
                  </td>
                  <td>
                    <Text>Dr. MV Panduranga Rao</Text>
                  </td>
                  <td>23.07.2018</td>
                </tr>
              );
            })}
          </tbody>
        </table>
        <Divider />
        List,
      </Box>
    </Flex>
  );
}

export default StudentView;
