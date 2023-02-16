import React, {useState, useEffect} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
import Layout from '../../components/Layout/Layout';
import InputForm from '../../components/Input-Form/Input-Form';
const CreatePost = (props) => {

// const [user, setUser] = useState();    

// const [isAuth, setIsAuth] = useState();
// const [isLoading, setIsLoading] = useState(true);

// useEffect(() => {
//   if (props.isAuth){

//     setIsAuth(props.setIsAuth)
//       setUser(props.username)
//       setIsLoading(false)
//       console.log("CreatePost:"+props.username);

//     }   
// }, [props.isAuth, props.username])  

    // return isLoading ? (

    //     <div>Loading...</div>
    //   ) : (
    //     <div>
    //     <Header />
    //     <Body createPost={true} username={user} isAuth={isAuth}/>
    //     </div>
    //   );
    console.log("CreatePOST:", props.username);
    return(
      <Layout>
        <InputForm/>
      </Layout>
    )
    
};

export default CreatePost;