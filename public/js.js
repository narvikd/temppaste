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
            if (data.success === false) {
                displayErr(data.message)
            } else {
                window.location.href= `/paste/${data.data}/raw`;
            }
        })
        .catch((e) => {
            displayErr(e)
        });
}

function displayErr(title) {
    Swal.fire({
        title: capitalize(title),
        icon: 'error',
        confirmButtonText: 'Ok'
    });
}

function capitalize(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
}
