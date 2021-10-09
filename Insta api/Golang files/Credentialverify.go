func DoSomeInstagramApiStuff(accessToken string) {
  api := New("", accessToken)

  if ok, err := api.VerifyCredentials(); !ok
  {
    panic(err)
  }
  
  var myId string
  if resp, err := api.GetSelf(); err != nil 
  {
    panic(err)
  } 
  else 
  {
    me := resp.User
    fmt.Printf("My UserID is %s and I have %d followers\n", me.Id, me.Counts.FollowedBy)
  }

  params := url.Values{}
  params.Set("count", "1")
  if resp, err := api.GetUserRecentMedia("self" /* this works :) */, params); err != nil
  {
    panic(err)
  } 
  else 
  {
    if len(resp.Medias) == 0 { // [sic]
      panic("Browse something in intagram you may have some posts to be seen!")
    }
    media := resp.Medias[0]
    fmt.Println("My last feed was %s with %d comments and %d likes. (url: %s)\n", media.Type, media.Comments.Count, media.Like.Count, media.Link)
  }
}
