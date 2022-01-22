import React from "react";
import ReactDOM from "react-dom";
import Form from "./configurator_form";

const onSubmt = ({ formData }) => alert("Data submitted: ", formData);

function App() {
//	return <Form onSubmit={onSubmit} />;
  return <Form onSubmit={values => alert(JSON.stringify(values))} />;
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);

