// script.js

const app = (function () {
    const game = {};
    const val1 = document.querySelector(".val1");
    const val2 = document.querySelector(".val2");
    const output = document.querySelector(".output");
    const answer = document.querySelector("input");

    function init() {
        document.querySelector("button").addEventListener("click", checker);
        loadQuestion();
    }

    function ranValue(min, max) {
        return Math.floor(Math.random() * (max - min + 1) + min);
    }

    function loadQuestion() {
        game.val1 = ranValue(1, 100);
        game.val2 = ranValue(1, 100);
        game.answer = game.val1 + game.val2;
        val1.textContent = game.val1;
        val2.textContent = game.val2;
    }

    function checker() {
        let bg = answer.value == game.answer ? "green" : "red";
        output.innerHTML += `<div style="font-size: 1.4em; font-weight: 600; color:${bg};">${game.val1} + ${game.val2} = ${game.answer}<br><br>Your answer: ${answer.value}<br><br><button class="play-again" onclick=location.reload()>Play again?</button></div>`;
        answer.value = "";
    }

    return {
        init: init,
    };
})();

document.addEventListener("DOMContentLoaded", app.init);
