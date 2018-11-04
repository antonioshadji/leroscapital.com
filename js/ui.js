/* global gapi */
/* eslint-disable no-unused-vars, camelcase */

function onSignIn (googleUser) {
  var profile = googleUser.getBasicProfile()
  console.log('ID: ' + profile.getId()) // Do not send to your backend! Use an ID token instead.
  console.log('Name: ' + profile.getName())
  console.log('Email: ' + profile.getEmail()) // This is null if the 'email' scope is not present.
  var id_token = googleUser.getAuthResponse().id_token
  console.log('ID Token for backend: ' + id_token)

  document.getElementById('t_loggedin_user').innerHTML =
  `${profile.getName()} <img class='rounded-circle' src='${profile.getImageUrl()}' height='36px'/>`
  document.getElementById('b_logout').classList.remove('d-none')
}

function signOut () {
  var auth2 = gapi.auth2.getAuthInstance()
  auth2.signOut().then(function () {
    console.log('User signed out.')
    document.getElementById('b_logout').classList.add('d-none')
    document.getElementById('t_loggedin_user').innerHTML = ''
  })
}

window.addEventListener('load', (event) => {
  /* utility code for all sites */
  document.getElementById('year').innerHTML = new Date().getFullYear()

  /* project specific */
  document.getElementById('b_logout').addEventListener('click', () => {
    console.log('click #b_logout')
    signOut()
  })
})
