res, err := api.GetUserRecentMedia("some-user", params)
if err != nil {
  panic(err)
}

done := make(chan bool) // creating a done channel when the feeds are seen
defer close(done)

medias, errs := api.IterateMedia(res, done)

for media := range medias {
  processMedia(media)

  if doneWithMedia(media) { //to exit quickly
    done <- true
  }
}

if err, ok := <- errs; ok && err != nil { // if moved out of insta by some kind of error then this is needed
  panic(err)
}
