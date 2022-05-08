function newPaste() {
    const paste = document.querySelector("textarea").value;
    fetch("/paste", {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },

        body: JSON.stringify({
            "content": paste,
        })
    })
        .then((res) => {
            return res.json()
        })
        .then((data) => {
            window.location.href= `/paste/${data.data}`;
        });
}
