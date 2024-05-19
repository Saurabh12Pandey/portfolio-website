"use client";

import React, { useEffect, useState } from "react";
import SectionHeading from "./section-heading";
import {
  VerticalTimeline,
  VerticalTimelineElement,
} from "react-vertical-timeline-component";
import "react-vertical-timeline-component/style.min.css";
import { experiencesData } from "@/lib/data";
import { useSectionInView } from "@/lib/hooks";
import { useTheme } from "@/context/theme-context";
import axios from "axios";
import { CgWorkAlt } from "react-icons/cg";
import { FaReact } from "react-icons/fa";
import { LuGraduationCap } from "react-icons/lu";
export default function Experience() {
  const { ref } = useSectionInView("Experience");
  const { theme } = useTheme();


  const [experience, setExperience] = useState([]);




  console.log(experiencesData);
  useEffect(() => {
    const getData = async () => {
      try {
        axios.get("http://localhost:8080/experience")
          .then((response) => {
            console.log(response)
            let experienceList = [];
            const { data } = response;
            data.map((exp, index) => {
              const { experienceDescription, experienceDuration, experienceTitle } = exp;
              experienceList.push({
                icon: index === 0 ? React.createElement(LuGraduationCap) : index === 1 ? React.createElement(CgWorkAlt) : React.createElement(FaReact),
                title: experienceTitle,
                location: "Remote",
                description: experienceDescription,
                date: experienceDuration,
              });
            });
            // setSkills(skills);
            // setHeaderText(about_me);
            setExperience(experienceList);
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
    <section id="experience" ref={ref} className="scroll-mt-28 mb-28 sm:mb-40">
      <SectionHeading>My experience</SectionHeading>
      <VerticalTimeline lineColor="">
        {experience && experience.length > 0 && experience.map((item, index) => (
          <React.Fragment key={index}>
            <VerticalTimelineElement
              contentStyle={{
                background:
                  theme === "light" ? "#f3f4f6" : "rgba(255, 255, 255, 0.05)",
                boxShadow: "none",
                border: "1px solid rgba(0, 0, 0, 0.05)",
                textAlign: "left",
                padding: "1.3rem 2rem",
              }}
              contentArrowStyle={{
                borderRight:
                  theme === "light"
                    ? "0.4rem solid #9ca3af"
                    : "0.4rem solid rgba(255, 255, 255, 0.5)",
              }}
              date={item.date}
              icon={item.icon}
              iconStyle={{
                background:
                  theme === "light" ? "white" : "rgba(255, 255, 255, 0.15)",
                fontSize: "1.5rem",
              }}
            >
              <h3 className="font-semibold capitalize">{item.title}</h3>
              <p className="font-normal !mt-0">{item.location}</p>
              <p className="!mt-1 !font-normal text-gray-700 dark:text-white/75">
                {item.description}
              </p>
            </VerticalTimelineElement>
          </React.Fragment>
        ))}
      </VerticalTimeline>
    </section>
  );
}
