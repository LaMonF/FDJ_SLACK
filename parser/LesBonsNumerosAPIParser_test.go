package parser

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Helpfull function to fake the api end point
func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}

	return cli, s.Close
}

func TestLesBonsNumerosAPIParser_GetLotteryResult(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(okResponse))
	})
	httpClient, teardown := testingHTTPClient(h)
	defer teardown()

	cli := NewParser()
	cli.httpClient = httpClient
	cli.url = "http://blabla.fr/"

	results, err := cli.GetLotteryResult()

	if err != nil {
		t.Fail()
	}

	if results.LuckyBall != 9 {
		t.Fail()
	}

	if !reflect.DeepEqual(results.Balls, []int{3, 40, 42, 45, 47}) {
		t.Fail()
	}

	if results.Date != "samedi  9 février 2019" {
		t.Fail()
	}

	if results.NextLotteryDate != "lundi 11 février 2019" {
		t.Fail()
	}

	if results.WinnerNumber != 0 {
		t.Fail()
	}

	if results.NextLotteryPrize != 12000000 {
		t.Fail()
	}

	if results.WinnerPrize != 0 {
		t.Fail()
	}
}

// STUB response from the API
const okResponse string = `
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/">
<channel>
<title>Résultats du Loto - LesBonsNumeros.com</title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/]]></link>
<description>Résultats et rapports du Loto (plus de nombreuses statistiques).</description>
<language>fr</language>
<item>
<title><![CDATA[Tirage Loto du samedi  9 février 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1624-samedi-9-fevrier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1624-samedi-9-fevrier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du samedi  9 février 2019</h1>
<h3>Numéros : 3 - 40 - 42 - 45 - 47</h3>
<h3>Numéro Chance : 9</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1624-samedi-9-fevrier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 47 - 40 - 45 - 42 - 3</p>
<p>Numéro Chance : 9</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du lundi 11 février 2019est de <strong>12 000 000 €</strong>.</p>
]]></description>
<pubDate>Sat, 09 Feb 2019 20:51:04 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du mercredi  6 février 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1623-mercredi-6-fevrier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1623-mercredi-6-fevrier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du mercredi  6 février 2019</h1>
<h3>Numéros : 4 - 15 - 21 - 22 - 26</h3>
<h3>Numéro Chance : 10</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1623-mercredi-6-fevrier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 4 - 15 - 21 - 26 - 22</p>
<p>Numéro Chance : 10</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du samedi  9 février 2019est de <strong>11 000 000 €</strong>.</p>
]]></description>
<pubDate>Wed, 06 Feb 2019 20:49:04 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du lundi  4 février 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1622-lundi-4-fevrier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1622-lundi-4-fevrier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du lundi  4 février 2019</h1>
<h3>Numéros : 12 - 24 - 37 - 40 - 42</h3>
<h3>Numéro Chance : 1</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1622-lundi-4-fevrier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 40 - 12 - 37 - 24 - 42</p>
<p>Numéro Chance : 1</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du mercredi  6 février 2019est de <strong>10 000 000 €</strong>.</p>
]]></description>
<pubDate>Mon, 04 Feb 2019 20:39:02 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du samedi  2 février 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1621-samedi-2-fevrier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1621-samedi-2-fevrier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du samedi  2 février 2019</h1>
<h3>Numéros : 5 - 7 - 27 - 32 - 35</h3>
<h3>Numéro Chance : 3</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1621-samedi-2-fevrier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 32 - 27 - 5 - 7 - 35</p>
<p>Numéro Chance : 3</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du lundi  4 février 2019est de <strong>9 000 000 €</strong>.</p>
]]></description>
<pubDate>Sat, 02 Feb 2019 20:47:03 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du mercredi 30 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1620-mercredi-30-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1620-mercredi-30-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du mercredi 30 janvier 2019</h1>
<h3>Numéros : 11 - 15 - 20 - 30 - 47</h3>
<h3>Numéro Chance : 9</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1620-mercredi-30-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 20 - 47 - 15 - 11 - 30</p>
<p>Numéro Chance : 9</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du samedi  2 février 2019est de <strong>8 000 000 €</strong>.</p>
]]></description>
<pubDate>Wed, 30 Jan 2019 20:51:03 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du lundi 28 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1619-lundi-28-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1619-lundi-28-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du lundi 28 janvier 2019</h1>
<h3>Numéros : 6 - 15 - 30 - 37 - 47</h3>
<h3>Numéro Chance : 6</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1619-lundi-28-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 37 - 15 - 30 - 6 - 47</p>
<p>Numéro Chance : 6</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du mercredi 30 janvier 2019est de <strong>7 000 000 €</strong>.</p>
]]></description>
<pubDate>Mon, 28 Jan 2019 20:49:04 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du samedi 26 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1618-samedi-26-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1618-samedi-26-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du samedi 26 janvier 2019</h1>
<h3>Numéros : 12 - 14 - 16 - 30 - 42</h3>
<h3>Numéro Chance : 2</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1618-samedi-26-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 30 - 14 - 12 - 16 - 42</p>
<p>Numéro Chance : 2</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du lundi 28 janvier 2019est de <strong>6 000 000 €</strong>.</p>
]]></description>
<pubDate>Sat, 26 Jan 2019 20:39:02 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du mercredi 23 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1617-mercredi-23-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1617-mercredi-23-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du mercredi 23 janvier 2019</h1>
<h3>Numéros : 17 - 24 - 25 - 38 - 41</h3>
<h3>Numéro Chance : 10</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1617-mercredi-23-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 25 - 17 - 41 - 24 - 38</p>
<p>Numéro Chance : 10</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du samedi 26 janvier 2019est de <strong>5 000 000 €</strong>.</p>
]]></description>
<pubDate>Wed, 23 Jan 2019 20:35:02 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du lundi 21 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1616-lundi-21-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1616-lundi-21-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du lundi 21 janvier 2019</h1>
<h3>Numéros : 17 - 20 - 26 - 30 - 47</h3>
<h3>Numéro Chance : 6</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1616-lundi-21-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 17 - 26 - 30 - 47 - 20</p>
<p>Numéro Chance : 6</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du mercredi 23 janvier 2019est de <strong>4 000 000 €</strong>.</p>
]]></description>
<pubDate>Mon, 21 Jan 2019 20:43:03 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du samedi 19 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1615-samedi-19-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1615-samedi-19-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du samedi 19 janvier 2019</h1>
<h3>Numéros : 10 - 16 - 34 - 37 - 39</h3>
<h3>Numéro Chance : 7</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1615-samedi-19-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 37 - 10 - 39 - 16 - 34</p>
<p>Numéro Chance : 7</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du lundi 21 janvier 2019est de <strong>3 000 000 €</strong>.</p>
]]></description>
<pubDate>Sat, 19 Jan 2019 20:41:03 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du mercredi 16 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1614-mercredi-16-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1614-mercredi-16-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du mercredi 16 janvier 2019</h1>
<h3>Numéros : 1 - 10 - 11 - 15 - 31</h3>
<h3>Numéro Chance : 3</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1614-mercredi-16-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 15 - 11 - 10 - 1 - 31</p>
<p>Numéro Chance : 3</p>
<h1>Jackpot</h1>
<p>
Un joueur a remporté le jackpot d'un montant de <strong>5 000 000&nbsp;€</strong>.</p><p>
Le montant du jackpot du prochain tirage du samedi 19 janvier 2019est de <strong>2 000 000 €</strong>.</p>
]]></description>
<pubDate>Wed, 16 Jan 2019 21:09:05 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du lundi 14 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1613-lundi-14-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1613-lundi-14-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du lundi 14 janvier 2019</h1>
<h3>Numéros : 1 - 11 - 33 - 35 - 44</h3>
<h3>Numéro Chance : 10</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1613-lundi-14-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 35 - 33 - 44 - 1 - 11</p>
<p>Numéro Chance : 10</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du mercredi 16 janvier 2019est de <strong>5 000 000 €</strong>.</p>
]]></description>
<pubDate>Mon, 14 Jan 2019 20:45:03 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du samedi 12 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1612-samedi-12-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1612-samedi-12-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du samedi 12 janvier 2019</h1>
<h3>Numéros : 11 - 19 - 24 - 25 - 34</h3>
<h3>Numéro Chance : 8</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1612-samedi-12-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 19 - 24 - 11 - 34 - 25</p>
<p>Numéro Chance : 8</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du lundi 14 janvier 2019est de <strong>4 000 000 €</strong>.</p>
]]></description>
<pubDate>Sat, 12 Jan 2019 20:45:03 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du mercredi  9 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1611-mercredi-9-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1611-mercredi-9-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du mercredi  9 janvier 2019</h1>
<h3>Numéros : 6 - 15 - 19 - 21 - 46</h3>
<h3>Numéro Chance : 10</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1611-mercredi-9-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 6 - 19 - 46 - 21 - 15</p>
<p>Numéro Chance : 10</p>
<h1>Jackpot</h1>
<p>
Le jackpot n'a pas été remporté lors de ce tirage !</p><p>
Le montant du jackpot du prochain tirage du samedi 12 janvier 2019est de <strong>3 000 000 €</strong>.</p>
]]></description>
<pubDate>Wed, 09 Jan 2019 20:55:04 +0100</pubDate>
</item>
<item>
<title><![CDATA[Tirage Loto du lundi  7 janvier 2019]]></title>
<link><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1610-lundi-7-janvier-2019.htm]]></link>
<guid><![CDATA[https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1610-lundi-7-janvier-2019.htm]]></guid>
<description><![CDATA[
<h1>Résultats du lundi  7 janvier 2019</h1>
<h3>Numéros : 1 - 7 - 12 - 19 - 31</h3>
<h3>Numéro Chance : 10</h3>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/rapports-tirage-1610-lundi-7-janvier-2019.htm">Voir les rapports complets</a></p>
<p><a href="https://www.lesbonsnumeros.com/loto/resultats/verifier-vos-gains.htm">Vérifiez vos gains</a></p>
<h1>Ordre de sortie</h1>
<p>Numéros : 1 - 19 - 12 - 31 - 7</p>
<p>Numéro Chance : 10</p>
<h1>Jackpot</h1>
<p>
Un joueur a remporté le jackpot d'un montant de <strong>3 000 000&nbsp;€</strong>.</p><p>
Le montant du jackpot du prochain tirage du mercredi  9 janvier 2019est de <strong>2 000 000 €</strong>.</p>
]]></description>
<pubDate>Mon, 07 Jan 2019 20:39:03 +0100</pubDate>
</item>
</channel>
</rss>
`
