import './style.css';
import './app.css';

import {BuildStudyPlan} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
  <main class="board">
    <section class="hero">
      <p class="eyebrow">04-go-frameworks</p>
      <h1>Go Framework Study Board</h1>
      <p class="subtitle">用一个桌面面板把 Web、RPC、CLI、ORM 和测试框架串起来，适合作为 Wails 入门示例。</p>
    </section>

    <section class="grid">
      <article class="card">
        <span class="tag">Web</span>
        <h2>Gin / Echo / Fiber / Chi</h2>
        <p>对比路由、中间件和 JSON API 写法，理解不同框架的心智模型。</p>
      </article>
      <article class="card">
        <span class="tag">RPC</span>
        <h2>gRPC</h2>
        <p>从 proto 文件出发，体验接口先行与代码生成的协作方式。</p>
      </article>
      <article class="card">
        <span class="tag">Tooling</span>
        <h2>Cobra / Testify</h2>
        <p>一个负责 CLI 组织，一个负责测试表达，都是工程里很常用的基础积木。</p>
      </article>
    </section>

    <section class="planner">
      <label class="label" for="name">输入你的名字，生成一条学习建议</label>
      <div class="inputRow">
        <input class="input" id="name" type="text" autocomplete="off" placeholder="比如：小王" />
        <button class="btn" id="build">生成建议</button>
      </div>
      <div class="result" id="result">这里会显示从 Go 后端返回的学习建议。</div>
    </section>
  </main>
`;

const nameElement = document.getElementById('name');
const resultElement = document.getElementById('result');
const buildButton = document.getElementById('build');

nameElement.focus();

window.buildPlan = function () {
    const name = nameElement.value;

    BuildStudyPlan(name)
        .then((result) => {
            resultElement.innerText = result;
        })
        .catch((err) => {
            console.error(err);
        });
};

buildButton.addEventListener('click', window.buildPlan);
nameElement.addEventListener('keydown', (event) => {
    if (event.key === 'Enter') {
        window.buildPlan();
    }
});
