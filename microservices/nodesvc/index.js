"use strict";

// load the express and morgan modules
const express = require("express");
const morgan = require("morgan");

const addr = process.env.ADDR || ":80";
const [host, port] = addr.split(":");
const portNum = parseInt(port);

const app = express();
app.use(morgan(process.env.LOG_FORMAT || "dev"));

app.get("/", (req, res) => {
	res.set("Content-Type", "text/plain");
	res.send("Hello, Node.js");
});

app.get("/v1/users/me/hello", (req, res) => {
	let userJSON = req.get("X-User");
	if (!userJSON) {
		throw new Error("No X-User header provided");
	}
	let user = JSON.parse(userJSON);
	res.json({
		message: `Hello, ${user.firstName} ${user.lastName}`
	});
});

// error handler (to prevent stacktrace being returned to client)
app.use((err, req, res, next) => {
	console.error(error.stack);
	// send only the error message back to the client
	res.set("Content-Type", "text/plain");
	res.send(err.message);
});

app.listen(portNum, host, () => {
	// template string syntax below
	console.log(`server is listening at http://${addr}...`); 
});
