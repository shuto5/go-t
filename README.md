# go-todo-project
シンプルなTodoアプリです。
# 技術スタック
・サーバー：Go<br>
 
・DB：SQLite<br>
    
・ページの表示：HTML<br>

# 操作概要
・トップページから新規タスク作成ボタンをクリックする。(index.html)<br>
・新規タスク作成画面からタイトルと詳細を入力して追加ボタンをクリックする。(create.html)<br>
・追加ボタンをクリックするとトップ画面に戻り、作成したTODOのタイトルが追加されており、横に編集、削除リンクもつくようにしている。<br>
・タイトルのリンクをクリックすると、詳細の内容が表示される。(detail.html)<br>
・編集リンクをクリックすると、編集画面に移行し、タイトルと詳細内容の編集ができる。(edit.html)<br>
・削除リンクをクリックすると、削除画面に移行し、はいボタンをクリックすると、作成したタイトルと詳細のTODOが削除される。(delete.html)<br>
![top](https://github.com/shuto5/go-todo-project/assets/85450386/66cad927-08a1-4f23-8282-804bea0cdf12)
![create](https://github.com/shuto5/go-todo-project/assets/85450386/c496eb08-e0e5-4686-8c1e-aebf98d24a61)<br>
トップ画面   →　新規タスク作成画面<br>


![createadd](https://github.com/shuto5/go-todo-project/assets/85450386/7a0fd16d-10d2-4d2c-bdde-3f20635c0e9a)
![topadd](https://github.com/shuto5/go-todo-project/assets/85450386/07e20baf-511d-44f2-851b-09c6aae973fb)<br>
タイトルと詳細内容を入力　　→　　追加ボタンをクリックするとタイトルが追加される<br>
![detail](https://github.com/shuto5/go-todo-project/assets/85450386/29085907-1e6b-41d4-9e3b-5fd1090f3606)<br>
タイトルリンクをクリック　→　詳細内容が表示される<br>
![edit](https://github.com/shuto5/go-todo-project/assets/85450386/d9474e76-5cb8-4855-a13a-cff65dd3d351)
![updatetop](https://github.com/shuto5/go-todo-project/assets/85450386/9138d5d5-2763-4b1f-bb6d-b9e8eff800c9)<br>
編集リンクをクリック　→　編集画面に移行　→　編集が完了したら更新ボタンをクリック　→　編集できているか確認する<br>
![delete](https://github.com/shuto5/go-todo-project/assets/85450386/89f77f66-e525-4d2f-a0fa-bc7df80a0f90)
![top](https://github.com/shuto5/go-todo-project/assets/85450386/66cad927-08a1-4f23-8282-804bea0cdf12)<br>
削除リンクをクリック →　はいボタンをクリック　→　削除が実行

# 工夫、苦戦した点
CRUD処理を施していることや、タイトルと詳細を登録するとき、詳細情報を多く記録できるように追加の際はタイトルだけ追加して、タイトルリンクをクリックすると詳細が見れるように設定したこと。

画面ごとにできる操作を変えている。一つ一つの画面でどういった操作ができるのかというのを明確にするため。


# 改善点、もう一度1から作るとしたいしたいこと
**Dockerを用いた環境からの開発**<br>
今回はVScodeからターミナルでgo getしてインポートしたが、Dockerを使うことで開発がしやすくなる。現在勉強中である。<br>
**GORMやフレームワークを用いた開発**<br>
GoのDB接続でGORMがあることを知ったのと、Goの理解を深めるためにあえてフレームワークは特に使わなかったが、アプリケーション開発ではフレームワークが必須になってくると考えており、それらを用いた開発もしてみたいと思った。<br>
**機能の追加やセキュリティに関して**<br>
今回は最低限の機能しか実装していないので、セキュリティを考慮したコードの実装をしてみたい。<br>
あとログイン機能やテストを書いたことがないので、そういったところも実装できるようにしたい。

