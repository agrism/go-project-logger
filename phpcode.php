<?php
/* start */

$name = 'xxx';

$name = (function () {
	$segments = explode('/', __DIR__);

	$segments = array_reverse($segments);

	array_shift($segments);

	return array_shift($segments);
})();

$name = $name.':';


$session = 'u_'.uniqid();
$path = ($_SERVER['HTTP_HOST'] ?? '').($_SERVER['REQUEST_URI'] ?? '');
$data = $_REQUEST;
$startTime = date('Y-m-d H:i:s');

$start = microtime(1);

send($name, $session, $path, null, $startTime, null, null);

register_shutdown_function(function () use ($start, $name, $startTime, $session, $path, $data) {
	$tabLength = 20;
	$end = microtime(1);
	$length = round($end - $start, 1);

	$content = [];

	$content[] = $name.str_repeat(' ', $tabLength - strlen($name));
	$content[] = '['.date('Y-m-d H:i:s', $start).' - '.date('Y-m-d H:i:s', $end).']';
	$content[] = ' ';
	$content[] = number_format($length, 3);
	$content[] = ' ';
	$content[] = $_SERVER['HTTP_HOST'] ?? '';
	$content[] = $_SERVER['REQUEST_URI'] ?? '';

	if (strpos($name, 'mobile') === 0) {
		$post = json_decode(file_get_contents('php://input'), 1);
		$content[] = PHP_EOL;

		$str = '* screen:';
		$content[] = $str.str_repeat(' ', $tabLength - strlen($str));
		$content[] = ($post['screen'] ?? 'n/a');
		$content[] = PHP_EOL;
		$str = '* action:';
		$content[] = $str.str_repeat(' ', $tabLength - strlen($str));
		$content[] = ($post['action'] ?? 'n/a');
		$content[] = PHP_EOL;
		$content[] = str_repeat('-', 50);
		$content[] = PHP_EOL;
	}

	$endTime = date('Y-m-d H:i:s');
	send($name, $session, $path, $data, $startTime, $endTime, $length);

	exec('cd ~ && echo "'.implode('', $content).'" >> exec_time.log');
});

function send($projectName, $session, $path, $data = [], $startTime, $endTime, $length)
{
	$ch = curl_init();
	$fields = [
		'project' => $projectName,
		'session' => $session,
		'path' => $path,
		'data' => $data,
		'start' => $startTime,
		'end' => $endTime,
		'length' => $length,
	];

	$postvars = $fields_string = http_build_query($fields);

	$url = 'http://laravel-test.test/log';
	curl_setopt($ch, CURLOPT_URL, $url);
	curl_setopt($ch, CURLOPT_POST, 1); //0 for a get request
	curl_setopt($ch, CURLOPT_POSTFIELDS, $postvars);
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
	curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 30);
	curl_setopt($ch, CURLOPT_TIMEOUT, 20);
	curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
	curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);

	curl_setopt($ch, CURLOPT_VERBOSE, true);
	$response = curl_exec($ch);

	if ($response === false) {
		file_put_contents('/Users/agris/exec_time.log', 'error:' . json_encode(curl_error($ch)), FILE_APPEND);
	}

	curl_close($ch);
}

/* end */