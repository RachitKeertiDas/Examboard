import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import {
  Flex,
  Box,
  Divider,
  Text,
  Stack,
  Heading,
  Button,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  List,
  ListItem,
  FormControl,
  FormLabel,
} from "@chakra-ui/core";
import "./styles.css";

function CourseRecord(props) {
  const { name, code } = props;
  const {
    isOpen: isListOpen,
    onOpen: openList,
    onClose: closeList,
  } = useDisclosure();
  const {
    isOpen: isExamDialogOpen,
    onOpen: openExamDialog,
    onClose: closeExamDialog,
  } = useDisclosure();
  const [StudentList, setStudentList] = useState([]);
  const [ExamList ,setExamList] = useState([]);
  return (
    <Flex
      shadow="md"
      className="course-list-box"
      justifyContent="space-between"
    >
      <Box>
        <Heading as="h4" size="md">
          {name}
        </Heading>
        <Text>{code}</Text>
      </Box>
      <Box className="course-button-container">
        <Button
          className="course-student-list-fetch-button"
          onClick={() => {
            fetch("http://localhost:8080/api/getCourseDetails/" + code)
              .then((response) => response.json())
              .then((responseBody) => {
                setStudentList(responseBody.Students);
              });
            openList();
          }}
        >
          See Students list
        </Button>
        <Modal isOpen={isListOpen} onClose={closeList} scrollBehavior="inside">
          <ModalOverlay />
          <ModalContent>
            <ModalHeader>Student List for Course {code}</ModalHeader>
            <ModalCloseButton />
            <ModalBody>
              <List>
                {StudentList.map((Student) => (
                 <ListItem key={Student}>{Student}</ListItem>
                ))
              }
              </List>
            </ModalBody>
            <ModalFooter>
              <Button onClick={closeList}>OK</Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
        <Button
          className="course-student-list-fetch-button"
          variantColor="teal"
          onClick={() => {
            fetch("http://localhost:8080/api/getCourseDetails"+code)
            .then((response) => response.json())
            .then((responseBody)=>{
              setExamList(responseBody.Exams);
            })
            openExamDialog();
          }}
        >
          Manage Exams
        </Button>
        <Modal isOpen={isExamDialogOpen} onClose={closeExamDialog}>
          <ModalOverlay />
          <ModalContent>
            <ModalHeader>Manage Exams for Course {code}</ModalHeader>
            <ModalCloseButton />
            <ModalBody>
              <List>
                <ListItem>CS19BTECH11034</ListItem>
                {ExamList.map((item) => (
                  <ListItem>{item}</ListItem>
                ))}
                <FormControl isRequired>
                  <FormLabel>Exam Date</FormLabel>
                  <FormLabel>Exam Time Range</FormLabel>
                </FormControl>
              </List>
            </ModalBody>
            <ModalFooter>
              <Button onClick={closeExamDialog}>OK</Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
      </Box>
    </Flex>
  );
}

CourseRecord.propTypes = {
  name: PropTypes.string.isRequired,
  code: PropTypes.string.isRequired,
};

const RecievedNames = [];
const RecievedCodes = [];

function InstructorView() {
  const [InstructorName, SetInstructorName] = useState(null);
  const [courseLength, setcourseLength] = useState(0);

  useEffect(() => {
    fetch("http://localhost:8080/api/instructor/getCourseList/")
      .then((res) => res.json())
      .then((responseBody) => {
        console.log("Response Body Recieved!");
        SetInstructorName(responseBody.Name);
        //Needs Refactoring
        console.log(RecievedNames.length);
        console.log(RecievedNames);
        for (let each in responseBody.Courses) {
          RecievedNames.push(responseBody.Courses[each]["Name"]);
        }

        for (let each in responseBody.Courses) {
          RecievedCodes.push(responseBody.Courses[each]["Code"]);
        }
        setcourseLength(
          Math.min(RecievedNames.length, 2 * responseBody.Courses.length)
        );
        console.log(courseLength); //0
      });
  }, [InstructorName, courseLength]);

  let CourseData = [];
  if (courseLength !== 2) {
    console.log(courseLength);
    for (let i = 0; i < courseLength / 2; i++) {
      let newObject = {};
      newObject["Name"] = RecievedNames[i];
      newObject["Code"] = RecievedCodes[i];
      CourseData.push(newObject);
    }
  }
  return (
    <Flex
      alignItems="center"
      flexDirection="column"
      justifyContent="space-between"
    >
      <Box className="instructor-home-heading">
        <Heading>Welcome, Dr. {InstructorName}!</Heading>
        <Divider />
      </Box>
      <Box>
        <Heading size="md">Your Courses</Heading>
        <Divider />
        <Stack>
          {CourseData.map((item) => (
            <CourseRecord name={item.Name} code={item.Code} />
          ))}
        </Stack>
      </Box>
    </Flex>
  );
}

export default InstructorView;

