import React from "react";
import InputForm from "../Input-Form/Input-Form";
import './Body.css'
import Post from "../Post/Post";
const Body = (props) => {  
return (
    <div className="body-container">
      <div className="body-posts-container">
      { props.createPost === true && <InputForm name= {props.name}/>}
      <Post />
      </div>
    {/* <InputForm name= {props.name}/> */}
    </div>
  )


};

export default Body;