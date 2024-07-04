'use client';

import Navbar from "../app/components/navBar";
import CoursesDisplay from "../app/components/CoursesDisplay.js";
import React from "react";

function Index() {

    return (
        <>
            <Navbar/>
            <h1 className="text-3xl mt-32 font-bold mb-4 text-center">课程列表</h1>
            <CoursesDisplay/>
        </>

    );
}

export default Index;
