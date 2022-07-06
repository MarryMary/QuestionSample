<?php
/*
 * 新規登録完了画面
 */
// 必要ファイルのインクルード
include dirname(__FILE__).'/../Tools/Session.php';
include dirname(__FILE__).'/../Tools/ValidateAndSecure.php';
include dirname(__FILE__).'/../Template/ServiceData.php';

// セッション開始
SessionStarter();

// 登録完了フラグがセッションに存在する場合
if(SessionIsIn('registration')){
    // セッション解除処理（ログアウト）
    SessionUnset();

    $title = 'Registration';
    $card_name = '新規登録';
    $errtype = False;

    // フォーム作成
    $form = <<<EOF
<p>登録が完了致しました。<br>
引き続きサービスを利用するにはログインして下さい。
</p>
<div style="text-align: center;">
    <button type="button" class="btn btn-primary" onclick="location.href='/{$SERVICE_ROOT}/Auth/login.php'" style='width: 90%'>ログイン</button>
</div>
EOF;

    // テンプレートファイル読み込み
    include dirname(__FILE__).'/../Template/BaseTemplate.php';
}else{
    // 登録完了時ではない場合
    header('Location: /'.$SERVICE_ROOT.'/Auth/login.php');
}