"use client";

import React, { useEffect, useState } from "react";
import SectionHeading from "./section-heading";
import { projectsData } from "@/lib/data";
import Project from "./project";
import { useSectionInView } from "@/lib/hooks";
import axios from "axios";

export default function Projects() {
  const { ref } = useSectionInView("Projects", 0.5);
  console.log(projectsData)
  const [projectDataForPage, setProjectData] = useState([]);
  

  useEffect(() => {
    const getData = async () => {
      try {
        axios.get("http://localhost:8080/projects")
          .then((response) => {
            // console.log(response)
            const { data } = response;
            let projectData = [];
            data.map((project) => {
              const { projectName, projectSummary,technologiesUsed, imagelink } = project;
              projectData.push({
                title: projectName,
                imageUrl: {
                  blurDataURL: "/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fcorpcomment.3895cd42.png&w=7&q=70",
                  blurHeight: 8,
                  blurWidth: 7,
                  height: 850,
                  src: imagelink,
                  width: 715,
                },
                tags: technologiesUsed,
                description: projectSummary,
              })
            });
            setProjectData(projectData);
            // setHeaderText(about_me);
          }).catch((error) => {
            console.log(error);
          })
      }
      catch (error) {
        alert(error);
      }
    };
    getData();
    
  }, []);


  return (
    <section ref={ref} id="projects" className="scroll-mt-28 mb-28">
      <SectionHeading>My projects</SectionHeading>
      <div>
        {projectDataForPage && projectDataForPage.length > 0 && projectDataForPage.map((project, index) => (
          <React.Fragment key={index}>
            <Project {...project} />
          </React.Fragment>
        ))}
      </div>
    </section>
  );
}
