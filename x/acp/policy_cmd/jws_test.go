package policy_cmd

// reject jws with critical header: If any of the listed extension Header Parameters are not understood
//and supported by the recipient, then the JWS is invalid
// reject jws with jose kid, x5c, x5u, kid, jwk, jku

// test jws containing kid and signed with kids key is rejected

// test did no verification method is rejected
