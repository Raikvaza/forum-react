import React, {useEffect, useState} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
import './HomePage.css'
import AuthContext from '../../index'
import Layout from '../../components/Layout/Layout';
import Post from '../../components/Post/Post';
const HomePage = (props) => {
const [posts, setPosts] = useState([]);
//const {isAuth} = useContext(AuthContext)

console.log(props);

useEffect(() => {
  console.log("POSTS FETCHED");
  const fetchData = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/home", {
        headers: {
          Accept: "application/json",
          Credentials: "include",
        },
        method: "GET",
        credentials: "include",
      });
      const data = await response.json();
      if (data !== null) {
        setPosts(data);
      } else {
        console.log("No posts in Homepage");
      }
    } catch (error) {
      console.error("Error fetching posts:", error);
      setPosts([]);
    }
  };
  fetchData();
}, []);

    console.log("POSTS:"+posts);
    
    const handlePosts = () => {
      if (!posts || posts.length === 0) {
        return <></>;
      }
    
      return posts.map(({PostId, Title, Content, CreationDate, Author})=>{
        return <Post key={PostId} postid={PostId} title={Title} content={Content} date={CreationDate} author={Author}/>
      })
    }
    return (
        
        <Layout>
          {handlePosts()}
        </Layout>
        // <div>
        // <Header status = {isAuth}/>
        // <Body createPost = {false} posts={posts} isAuth={isAuth}/>
        // <div>
        // </div>
        // </div>
    );
};

export default HomePage;