/*
 * Node.js server backend to handle AJAX calls 
 * and manage files
 */
var http = require('http');
var url = require('url');
var fs = require('fs');

const port = 8080

http.createServer(function (req, res) {

  var parsed = url.parse(req.url, true);
switch (parsed.pathname) {
// user
        case "/read":
  var filename = parsed.query['schema'];
  fs.readFile(filename, function(err, data) {
    if (err) {
      res.writeHead(404, {'Content-Type': 'application/json'});
      return res.end("404 Not Found");
    }  
      res.writeHead(200, {'Content-Type': 'application/json'});
    res.write(data);
    return res.end();
  });
            break
	case "/write":
		if (req.method === "POST" ) {
		var body = [];

  var filename = parsed.query['name'];

        req.on('data', function(chunk) {
            body.push(chunk);
        }).on('end', function() {
            data = Buffer.concat(body).toString();
            if (data) {
		    fs.writeFile(filename, data, 'utf-8', function (err) {
if (err) return console.log(err);
		    });
            res.end('File saved');
	    }
        });
			//fs.writeFile(filename, 

		} else {
      res.writeHead(400, {'Content-Type': 'application/json'});
      return res.end("Write operations require POST");
		}
            break

	default:
      res.writeHead(404, {'Content-Type': 'application/json'});
      return res.end("404 Not Found");
    }
}).listen(port);

