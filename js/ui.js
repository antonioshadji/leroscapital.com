/* global gapi */
/* eslint-disable no-unused-vars */

function onSignIn (googleUser) {
  var profile = googleUser.getBasicProfile()
  console.log('ID: ' + profile.getId()) // Do not send to your backend! Use an ID token instead.
  console.log('Name: ' + profile.getName())
  console.log('Image URL: ' + profile.getImageUrl())
  console.log('Email: ' + profile.getEmail()) // This is null if the 'email' scope is not present.
  document.getElementById('t_loggedin_user').innerHTML = profile.getName()
  document.getElementById('b_logout').classList.remove('d-none')
}

function signOut () {
  var auth2 = gapi.auth2.getAuthInstance()
  auth2.signOut().then(function () {
    console.log('User signed out.')
  })
}

window.addEventListener('load', (event) => {
  /* utility code for all sites */
  document.getElementById('year').innerHTML = new Date().getFullYear()

  /* project specific */
  document.getElementById('b_logout').addEventListener('click', () => {
    console.log('click #b_login')
    document.getElementById('b_logout').classList.add('d-none')
    document.getElementById('t_loggedin_user').innerHTML = ''
    signOut()
  })
})
