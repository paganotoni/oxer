package webpack

import (
	"io/ioutil"
	"os"
	"oxer/internal/meta"
	"path/filepath"
	"strings"
	"testing"
)

// The classic webpack.config.js
var classic = `
const Webpack = require("webpack");
const Glob = require("glob");
const path = require("path");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const ManifestPlugin = require("webpack-manifest-plugin");
const CleanObsoleteChunks = require('webpack-clean-obsolete-chunks');
const UglifyJsPlugin = require("uglifyjs-webpack-plugin");
const LiveReloadPlugin = require('webpack-livereload-plugin');

const configurator = {
	entries: function(){
		var entries = {
		application: [
			'./node_modules/jquery-ujs/src/rails.js',
			'./assets/css/application.scss',
		],
		}

		Glob.sync("./assets/*/*.*").forEach((entry) => {
		if (entry === './assets/css/application.scss') {
			return
		}

		let key = entry.replace(/(\.\/assets\/(src|js|css|go)\/)|\.(ts|js|s[ac]ss|go)/g, '')
		if(key.startsWith("_") || (/(ts|js|s[ac]ss|go)$/i).test(entry) == false) {
			return
		}

		if( entries[key] == null) {
			entries[key] = [entry]
			return
		}

		entries[key].push(entry)
		})
		return entries
	},

	plugins() {
		var plugins = [
		new CleanObsoleteChunks(),
		new Webpack.ProvidePlugin({$: "jquery",jQuery: "jquery"}),
		new MiniCssExtractPlugin({filename: "[name].[contenthash].css"}),
		new CopyWebpackPlugin([{from: "./assets",to: ""}], {copyUnmodified: true,ignore: ["css/**", "js/**", "src/**"] }),
		new Webpack.LoaderOptionsPlugin({minimize: true,debug: false}),
		new ManifestPlugin({fileName: "manifest.json"})
		];

		return plugins
	},

	moduleOptions: function() {
		return {
		rules: [
			{
			test: /\.s[ac]ss$/,
			use: [
				MiniCssExtractPlugin.loader,
				{ loader: "css-loader", options: {sourceMap: true}},
				{ loader: "sass-loader", options: {sourceMap: true}}
			]
			},
			{ test: /\.tsx?$/, use: "ts-loader", exclude: /node_modules/},
			{ test: /\.jsx?$/,loader: "babel-loader",exclude: /node_modules/ },
			{ test: /\.(woff|woff2|ttf|svg)(\?v=\d+\.\d+\.\d+)?$/,use: "url-loader"},
			{ test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,use: "file-loader" },
			{ test: require.resolve("jquery"),use: "expose-loader?jQuery!expose-loader?$"},
			{ test: /\.go$/, use: "gopherjs-loader"}
		]
		}
	},

	buildConfig: function(){
		const env = process.env.NODE_ENV || "development";

		var config = {
		mode: env,
		entry: configurator.entries(),
		output: {filename: "[name].[hash].js", path: ` + "${__dirname}/public/assets" + `},
		plugins: configurator.plugins(),
		module: configurator.moduleOptions(),
			resolve: {
				extensions: ['.ts', '.js', '.json']
			}
		}

		if( env === "development" ){
			config.plugins.push(new LiveReloadPlugin({appendScriptTag: true}))
			return config
		}

		const uglifier = new UglifyJsPlugin({
			uglifyOptions: {
				beautify: false,
				mangle: {keep_fnames: true},
				output: {comments: false},
				compress: {}
			}
		})

		config.optimization = {
			minimizer: [uglifier]
		}

		return config
	}
}

module.exports = configurator.buildConfig()
`

func TestReplaceAssetsRoutes(t *testing.T) {
	d := t.TempDir()
	os.Chdir(d)

	err := ioutil.WriteFile(filepath.Join(d, "webpack.config.js"), []byte(classic), 0777)
	if err != nil {
		t.Error(err)
	}

	err = ReplaceAssetsRoutes(meta.Application{})
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(d, "webpack.config.js"))
	if err != nil {
		t.Error(err)
	}

	shouldHaveChanged := []string{
		`new CopyWebpackPlugin([{from: "./assets"`,
		`if (entry === './assets`,
		`'./assets/css/application.scss'`,
		`entry.replace(/(\.\/assets\/(src`,
	}

	for _, change := range shouldHaveChanged {
		if strings.Contains(string(b), change) {
			t.Errorf("found an unfixed one")
		}
	}

}
