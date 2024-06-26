"use client";

import React, { useEffect, useState } from "react";
import SectionHeading from "./section-heading";
import { motion } from "framer-motion";
import { useSectionInView } from "@/lib/hooks";
import axios from "axios";

export default function About() {
  const { ref } = useSectionInView("About");

  const [headerText, setHeaderText] = useState("");
  

  useEffect(() => {
    const getData = async () => {
      try {
        axios.get("http://localhost:8080/home")
          .then((response) => {
            console.log(response)
            const { data } = response;
            const { about_me } = data[0];

            setHeaderText(about_me);
          }).catch((error) => {
            console.log(error);
          })
      }
      catch (error) {
        alert(error);
      }
    };
    getData();
    
  }, [])

  return (
    <motion.section
      ref={ref}
      className="mb-28 max-w-[45rem] text-center leading-8 sm:mb-40 scroll-mt-28"
      initial={{ opacity: 0, y: 100 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ delay: 0.175 }}
      id="about"
    >
      <SectionHeading>About me</SectionHeading>
      <p className="mb-3">
        {headerText}
      </p>

      <p>
        <span className="italic">When I'm not coding</span>, I enjoy playing
        video games, watching movies, and playing with my dog. I also enjoy{" "}
        <span className="font-medium">learning new things</span>. I am currently
        learning about{" "}
        <span className="font-medium">history and philosophy</span>. I'm also
        learning how to play the guitar.
      </p>
    </motion.section>
  );
}
