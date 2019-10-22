import * as React from "react";
import * as ReactDOM from "react-dom";
import axios from "axios";

axios
  .get("http://localhost:8080")
  .then(function(response) {
    // handle success
    console.log(response);
  })
  .catch(function(error) {
    // handle error
    console.log(error);
  })
  .finally(function() {
    console.log("Done");
  });

ReactDOM.render(<div>Hello World</div>, document.getElementById("root"));
