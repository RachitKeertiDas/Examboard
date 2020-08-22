import React from "react";
import { FaGoogle } from "react-icons/fa";
import { Flex, Heading, Box, Divider, Button } from "@chakra-ui/core";
import "./styles.css";

function Landing() {
  return (
    <Flex
      flexDirection="column"
      alignItems="center"
      justifyContent="space-between"
    >
      <Box>
        <Heading className="landing-heading">EXAM BOARD</Heading>
        <Divider />
      </Box>
      <Button leftIcon={FaGoogle}>Sign In With Google</Button>
    </Flex>
  );
}

export default Landing;
