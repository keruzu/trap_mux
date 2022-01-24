import React from "react";
import ReactDOM from "react-dom";
import Form from "./configurator_form";

const saveData = ({ formData }) => alert("Data submitted: ", formData);

function App() {
	return <Form onSubmit={saveData} />;
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);

