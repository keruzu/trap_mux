import React from "react";
import JSONSchemaForm from "@rjsf/core";
import "bootstrap/dist/css/bootstrap.css";

import postSchema from './schema/trapmux.json';

//const schemaUrl = "http://localhost:8080/read?schema=src/schema/trapmux.json"

/*
var postSchema = {}

fetch(schemaUrl)
  .then(response => response.json())
  .then((jsonData) => {
	  postSchema = jsonData
	  return jsonData
  })
  .catch((error) => {
    // handle your errors here
    console.error(error);
	  return {}
  })

class ConfigSelector extends React.Component {
  constructor(props) {
    super(props);
    this.state = {name: 'trapmux', config: {}, schema: {}}

// This binding is necessary to make `this` work in the callback
    this.handleClick = this. LoadSchemaAndConfig.bind(this);
  }

  componentDidMount() {
      this.state.config = LoadSchemaAndConfig(this.state.name)
  }
 LoadSchemaAndConfig() {
    this.setState({
      config: LoadSchemaAndConfig(this.name)
    });
  }

  render() {
  }
}
*/

export default function Form({ onSubmit }) {
  return (
    <div className="container">
      <div className="row">
        <div className="col-md-6">
          <JSONSchemaForm onSubmit={onSubmit} schema={postSchema} />
        </div>
      </div>
    </div>
  );
}

