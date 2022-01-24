import React from "react";
import JSONSchemaForm from "@rjsf/core";
import "bootstrap/dist/css/bootstrap.css";

import postSchema from './schema.json';

export default function Form({ onSubmit }) {
  return (
    <div class="container">
      <div class="row">
        <div class="col-md-6">
          <JSONSchemaForm onSubmit={onSubmit} schema={postSchema} />
        </div>
      </div>
    </div>
  );
}

