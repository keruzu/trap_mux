import React from "react";
import ReactDOM from "react-dom";
import axios from "axios"
import Form from "./configurator_form";

//const saveData = ({ formData }) => alert("Data submitted: ", formData);

//const baseUrl = "http://localhost:8080/write?name=src/data/trapmux.json"
  function saveData(formData) {
    axios
      .post("http://localhost:8080/save/trapmux", formData, {
headers: { 'Content-Type': 'application/json', 
     'Access-Control-Allow-Origin': '*'
})
      .then((response) => {
        setPost(response.data);
      });
  }

function App() {
	return <Form onSubmit={saveData} />;
}

const rootElement = document.getElementById("root");
ReactDOM.render(<App />, rootElement);

