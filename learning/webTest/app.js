async function getData() {
  const root = document.getElementById("root");
  const apiURL = "https://nuptials-accuracy-anyway.ngrok-free.dev/games";
  // const apiURL = "http://localhost:8080/games";

  if (!root) {
    return;
  }

  root.textContent = "Loading...";

  try {
    const resp = await fetch(apiURL, {
      headers: {
        "ngrok-skip-browser-warning": "true",
      },
    });

    if (!resp.ok) {
      throw new Error(`Request failed: ${resp.status}`);
    }

    const data = await resp.json();
    root.textContent = JSON.stringify(data, null, 2);
  } catch (err) {
    root.textContent = `Error: ${err.message}`;
  }
}

document.addEventListener("DOMContentLoaded", getData);
