import React from "react";

const Login = () => {
  return (
    <div>
      <div class="mx-auto flex justify-center py-10">
        <strong class="text-5xl p-5 tracking-widest shadow-xl">
          学生成绩管理系统
        </strong>
      </div>
      <div class="mx-auto my-19 flex h-[300px] w-[350px] flex-col border-2 bg-white text-black shadow-xl">
        <div class="mx-8 mt-7 mb-1 flex flex-row justify-start space-x-2">
          <div class="h-7 w-3 bg-[#0DE6AC]"></div>
          <div class="w-3 text-center font-sans text-xl font-bold">
            <h1>Login</h1>
          </div>
        </div>
        <div class="flex flex-col items-center">
          <input
            class="my-2 w-72 border p-2"
            type="email"
            placeholder="Username"
          />
          <input
            class="my-2 w-72 border p-2"
            type="password"
            placeholder="Password"
          />
        </div>

        <div class="flex justify-between text-sm mx-8 p-3">
          <label for="remember">
            <input type="radio" id="remember" name="a" checked="checked" />{" "}
            Student
          </label>
          <label for="remember">
            <input type="radio" id="remember" name="a" /> Teacher
          </label>
          <label for="remember">
            <input type="radio" id="remember" name="a" /> Administrator
          </label>
        </div>
        <div class="my-2 flex justify-center">
          <button class="w-72 border bg-[#0DE6AC] p-2 font-sans">Login</button>
        </div>
      </div>
    </div>
  );
};

export default Login;
