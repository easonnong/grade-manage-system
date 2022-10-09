import React from "react";

function consoleLog() {
  console.log("studentInfo");
}
const studentInfo = () => {
  return <div className="">{consoleLog()}</div>;
};

export default studentInfo;
