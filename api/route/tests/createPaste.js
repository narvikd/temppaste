pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response time is less than 200ms", function () {
    pm.expect(pm.response.responseTime).to.be.below(200);
});

pm.test("Response success", function () {
    const resp = pm.response.json();
    pm.expect(resp).to.have.ownProperty("success");
    pm.expect(resp.success).to.equal(true);
})

pm.test("Response includes data not empty", function () {
    const resp = pm.response.json();
    pm.expect(resp).to.have.ownProperty("data");
    pm.expect(resp.data).to.not.equal("");
})

pm.test("Check if paste can be retrieved", function () {
    const resp = pm.response.json();
    // Check if data exists and it's not empty
    pm.expect(resp).to.have.ownProperty("data");
    pm.expect(resp.data).to.not.equal("");

    const pasteID = resp.data;
    const endpoint = "http://127.0.0.1:3001/" + "paste/" + pasteID;

    pm.sendRequest(endpoint, function (err, response) {
        const resp2 = response.json();
        // Check if data exists and it's not empty
        pm.expect(resp2).to.have.ownProperty("data");
        pm.expect(resp2.data).to.not.equal("");
        // Check if id exists, it's not empty and it's equal to the post-paste ID
        pm.expect(resp2.data).to.have.ownProperty("id");
        pm.expect(resp2.data.id).to.not.equal("");
        pm.expect(resp2.data.id).to.equal(pasteID);
    });
})
