const posts = [
  {
    name: "Jimy",
    nickname: "merlins",
    content: "Hello World",
    reply: true,
    retweet: true,
    like: true,
    share: "https://twitter.com/post/21",
  },
  {
    name: "Jimy",
    nickname: "merlins",
    content: "nothing better than c",
    reply: 123567,
    retweet: 8678,
    like: 41,
    share: "https://twitter.com/post/22",
  },
  {
    name: "Jimy",
    nickname: "merlins",
    content: "stay hungry, stay foolish",
    reply: 475,
    retweet: 5123,
    like: 312,
    share: "https://twitter.com/post/23",
  },
  {
    name: "Jimy",
    nickname: "merlins",
    content:
      "Programming isn't about what you know; it's about what you can figure out.",
    reply: 51124,
    retweet: 123,
    like: 5142,
    share: "https://twitter.com/post/24",
  },
];

const listPost = (newPost) => {
  posts.map((post) => {
    console.log(post.nickname, post.content);
  });
};

const createNewPost = (newPost) => {
  const prom = new Promise((resolve, reject) => {
    posts.push(newPost);
    resolve(posts);
    // reject("Error is post not add to posts list.")
  });
  return prom;
};

async function showPosts() {
  try {
    await createNewPost({
      name: "Jimy",
      nickname: "merlins",
      content: "The best error message is the one that never shows up.",
      reply: 5124,
      retweet: 1123,
      like: 59142,
      share: "https://twitter.com/post/25",
    });
    listPost();
  } catch (error) {
    console.log(error);
  }
}

showPosts();
