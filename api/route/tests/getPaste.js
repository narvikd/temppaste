pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response time is less than 200ms", function () {
    pm.expect(pm.response.responseTime).to.be.below(200);
});

pm.test("Response includes data not empty", function () {
    const resp = pm.response.json();
    pm.expect(resp).to.have.ownProperty("data");
    pm.expect(resp.data).to.not.equal("");
})
