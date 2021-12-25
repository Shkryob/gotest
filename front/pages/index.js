import Head from 'next/head'
import { useState, useEffect } from 'react';
import gotestApi from '../services/GotestAPI'

export default function Home() {

  const [data, setData] = useState({ users: [], error: '' });
  useEffect(async () => {
    const result = await gotestApi.getList();

    setData((prevState) => {
      return {
        ...prevState,
        users: result.data.users,
      }
    });
  }, []);

  const createNewUser = (event) => {
    event.preventDefault();

    const data = new FormData(event.currentTarget);
    const formObject = Object.fromEntries(data.entries());

    gotestApi.create({user: formObject}).then((response) => {
      setData((prevState) => {
        return {
          ...prevState,
          users: [...prevState.users, response.data],
          error: '',
        }
      });
    }).catch((result) => {
      console.log('catch result', result);
      setData((prevState) => {
        return {
          ...prevState,
          error: result.response.data.errors.body,
        }
      });
    });
  };

  return (
    <div>
      <Head>
        <title>GoTest</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1>
          Welcome to the Club!
        </h1>

        <h2>New member</h2>
        <form onSubmit={createNewUser}>
          <label>
            Name: <input type="text" name="username" required />
          </label>
          <label>
            Email: <input type="email" name="email" required />
          </label>

          <div>
            <button type="submit">Add</button>
            <button type="reset">Clear</button>
          </div>
          <div className="error">{data.error}</div>
        </form>

        <h2>Members</h2>
        <table>
          <thead>
            <tr>
              <th>#</th>
              <th>Name</th>
              <th>Email</th>
              <th>Registration Date</th>
            </tr>
          </thead>
          <tbody>
             {data.users.map(({ user }) => (
              <tr key={user.id}>
                <td>{user.id}</td>
                <td>{user.username}</td>
                <td>{user.email}</td>
                <td>{user.createdAt}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </main>

      <style jsx>{`
        label {display: block;}
        .error {color: red;}
      `}</style>
    </div>
  )
}
