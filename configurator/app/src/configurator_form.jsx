import React from "react";
import JSONSchemaForm from "react-jsonschema-form";
import "bootstrap/dist/css/bootstrap.css";

import postSchema from './schema.json';
import postUiSchema from './uischema.json';

export default function Form({ onSubmit }) {
  return (
    <div class="container">
      <div class="row">
        <div class="col-md-6">
          <JSONSchemaForm onSubmit={onSubmit} schema={postSchema} uischema={postUiSchema} />
        </div>
      </div>
    </div>
  );
}

