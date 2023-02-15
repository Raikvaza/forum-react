import React, {useEffect, useState} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
import './HomePage.css'
// import { json } from 'react-router-dom';
const HomePage = (props) => {
const [posts, setPosts] = useState([]);
// const [fetchCalled, setFetchCalled] = useState(false);
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
    // if (!fetchCalled){
    //   return <div>LOADING</div>
    // }
    return (
        <div>
        <Header status = {props.isAuth}/>
        <Body createPost = {false} posts={posts} isAuth={props.isAuth}/>
        <div>
        </div>
        </div>
    );
};

export default HomePage;