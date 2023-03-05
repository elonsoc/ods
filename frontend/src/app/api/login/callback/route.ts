// this route is called by the user after a supposedly successful
// login attempt. The user is redirected to this route by the
// external IdP and the IdP sends a token as a query parameter.
// We need to extract the token and send it to the backend
// to verify it with the partner and get the user's data
// and a durable token to store so we don't have to login again.

// Also this all happens serverside  so we can pass secrets :-)

export async function GET(request: Request) {
  // for some reason we have
  if (request.url.indexOf("?") === -1) {
    return new Response("No token provided", { status: 400 });
  }

  // extract the token from the query parameters
  const token = request.url.split("?")[1].split("=")[1];

  if (token === null) {
    return new Response("No token provided", { status: 400 });
  }

  console.log("token", token);

  fetch("http://localhost:1337/identity/validate", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ token }),
  })
    .then((res) => {
      console.log(res.status);
      console.log(res.url);
    })
    .catch((err) => {
      console.log(err);
      return new Response(err, { status: 500 });
    });

  let response = new Response("OK", { status: 200 });
  response.headers.set("Set-Cookie", "token=" + token);
  return response;
}
