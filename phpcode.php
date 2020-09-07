<?php
/* start */

$name = 'xxx';

$start = microtime(1);
register_shutdown_function(function () use ($start, $name) {
	$tabLength = 20;
	$end = microtime(1);
	$length = round($end - $start, 1);

	$name = (function () {
		$segments = explode('/', __DIR__);

		$segments = array_reverse($segments);

		array_shift($segments);

		return array_shift($segments);
	})();

	$name = $name . ':';

    $content = [];
    
	$content[] = $name . str_repeat(' ', $tabLength - strlen($name));
	$content[] = '[' . date('Y-m-d H:i:s', $start) . ' - ' . date('Y-m-d H:i:s', $end) . ']';
	$content[] = ' ';
	$content[] = number_format($length, 3);

	if (strpos($name, 'mobile') === 0) {
		$post = json_decode(file_get_contents('php://input'), 1);
		$content[] = PHP_EOL;

		$str = '* screen:';
		$content[] = $str . str_repeat(' ', $tabLength - strlen($str));
		$content[] = ($post['screen'] ?? 'n/a');
		$content[] = PHP_EOL;
		$str = '* action:';
		$content[] = $str . str_repeat(' ', $tabLength - strlen($str));
		$content[] = ($post['action'] ?? 'n/a');
		$content[] = PHP_EOL;
		$content[] = str_repeat('-', 50);
		$content[] = PHP_EOL;
	}

	exec('cd ~ && echo "' . implode('', $content) . '" >> exec_time.log');
});

/* end */